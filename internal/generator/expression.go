package generator

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/woven-planet/go-zserio/internal/ast"
	"github.com/woven-planet/go-zserio/internal/parser"
)

func IdentifierToGoString(scope ast.Scope, expression *ast.Expression) string {
	dummyType := &ast.TypeReference{
		Name:    expression.ResultSymbol.Name,
		Package: expression.ResultSymbol.Package,
	}
	name, err := GoType(scope, dummyType)
	if err != nil {
		return "UNKNOWN_TYPE"
	}

	switch n := expression.ResultSymbol.Symbol.(type) {
	case *ast.Field:
		return "v." + n.Name
	case *ast.Parameter:
		return "v." + n.Name
	case *ast.Enum:
		return name
	case *ast.BitmaskType:
		return name
	case *ast.Function:
		return "v." + n.Name
	case *ast.Subtype:
		return name
	default:
		return "UNSUPPORTED_TYPE"
	}
}

// bracketedExpressionToGoString prints an array access expression,
// such as array[index].
func bracketedExpressionToGoString(scope ast.Scope, expression *ast.Expression) string {
	return fmt.Sprintf("%s[%s]",
		ExpressionToGoString(scope, expression.Operand1),
		ExpressionToGoString(scope, expression.Operand2),
	)
}

func parenthesizedExpressionToGoString(scope ast.Scope, expression *ast.Expression) string {
	return fmt.Sprintf("(%s)", ExpressionToGoString(scope, expression.Operand1))
}

func functionCallExpressionToString(scope ast.Scope, expression *ast.Expression) string {
	return fmt.Sprintf("%s()", ExpressionToGoString(scope, expression.Operand1))
}

func dotOperatorToGoString(scope ast.Scope, expression *ast.Expression) string {
	leftText := ExpressionToGoString(scope, expression.Operand1)
	rightText := strings.Title(expression.Operand2.Text)

	// Enums don't exist in Go, so the dot expression must be translated to the
	// constant, which is Enum name, followed by the Enum value.
	// The same is valid for bitmasks.
	if expression.Operand1.ResultType == ast.ExpressionTypeEnum {
		// Due to Go not supporting Enums, we need to check if subtyping is used.
		// Instead of generating the Enum value with the subtyped name, we need the
		// original type name.
		// For example, instead of:
		// SubTypeColorBlue
		// we want:
		// ColorBlue
		// because only the latter one is defined in the generated Go code.
		// First, resolve the original type from the scope.
		originalType, err := scope.OriginalType(
			&ast.TypeReference{
				Name:    expression.Operand1.ResultSymbol.Name,
				Package: expression.Operand1.ResultSymbol.Package,
			})
		if err != nil {
			return "ENUM_TYPE_LOOKUP_FAILED"
		}
		// Generate the enum value using the original type name, as well as the
		// enumeration value.
		originalEnumType, err := GoType(scope, originalType.Type)
		if err != nil {
			return "ENUM_GO_TYPE_DEDUCTION_FAILED"
		}
		enumValue := originalEnumType + strcase.ToCamel(strings.ToLower(expression.Operand2.Text))
		if originalType.RequiresCast {
			enumValue = fmt.Sprintf("%s(%s)", leftText, enumValue)
		}
		return enumValue
	} else if expression.Operand1.ResultType == ast.ExpressionTypeBitmask {
		return leftText + rightText
	}
	return fmt.Sprintf("%s.%s", leftText, rightText)
}

func numBitsOperatorToGoString(scope ast.Scope, expression *ast.Expression) string {
	return fmt.Sprintf("ztype.NumBits(uint64(%s))", ExpressionToGoString(scope, expression.Operand1))
}

// twoOperatorEqualTypesToGoString prints out an expression that uses two
// operands, where the operator expects to be both sides of the same type.
// since there are many different integer and type to integer types, casting
// might be necessary here.
func twoOperatorEqualTypesToGoString(scope ast.Scope, expression *ast.Expression) string {
	operator := ""
	switch expression.Type {
	case parser.ZserioParserPLUS:
		operator = "+"
	case parser.ZserioParserMINUS:
		operator = "-"
	case parser.ZserioParserMULTIPLY:
		operator = "*"
	case parser.ZserioParserDIVIDE:
		operator = "/"
	case parser.ZserioParserMODULO:
		operator = "%"
	case parser.ZserioParserLT:
		operator = "<"
	case parser.ZserioParserLE:
		operator = "<="
	case parser.ZserioParserGT:
		operator = ">"
	case parser.ZserioParserGE:
		operator = ">="
	case parser.ZserioParserEQ:
		operator = "=="
	case parser.ZserioParserNE:
		operator = "!="
	}

	operand1Str := ExpressionToGoString(scope, expression.Operand1)
	operand2Str := ExpressionToGoString(scope, expression.Operand2)

	// Check if casting is needed because of different original types.
	// Casting is needed in the following cases:
	// - mixing floats with integers
	// - mixing signed/unsigned types
	// - mixing different bit lengths
	// - mixing numerals of different types, e.g. 17 + 3.5
	// Casting is not needed in the following cases:
	// - The expression result is not an integer.
	// - One or both operands are numerals without type annotations, for example
	//   "17.5 + 13.6"
	//   In that case, NativeZserioType will be set to nil. Both operands must be
	//   integer or float numerals.
	if (expression.ResultType == ast.ExpressionTypeInteger || expression.ResultType == ast.ExpressionTypeFloat) && expression.NativeZserioType != nil {

		// Retrieve the resulting type of the expression
		expressionGoType, err := GoType(scope, expression.NativeZserioType)
		if err != nil {
			panic(err)
		}
		// Check if operand1 or operand2 (or both) needs a type cast.
		// In very rare cases, both operands need to be casted:
		// For example, mixing an uint8 with an int32 will require both
		// operands to be casted to uint32 (highest rank + unsigned wins).
		operand1RequiresCast := false
		if expression.Operand1.ResultType != expression.ResultType {
			operand1RequiresCast = true
		} else if expression.Operand1.NativeZserioType != nil {
			operand1GoType, err := GoType(scope, expression.Operand1.NativeZserioType)
			if err != nil {
				panic(err)
			}
			if operand1GoType != expressionGoType {
				operand1RequiresCast = true
			}
		}
		operand2RequiresCast := false
		if expression.Operand2.ResultType != expression.ResultType {
			operand2RequiresCast = true
		} else if expression.Operand2.NativeZserioType != nil {
			operand2GoType, err := GoType(scope, expression.Operand2.NativeZserioType)
			if err != nil {
				panic(err)
			}
			if operand2GoType != expressionGoType {
				operand2RequiresCast = true
			}
		}

		if operand1RequiresCast {
			operand1Str = fmt.Sprintf("%s(%s)", expressionGoType, operand1Str)
		}
		if operand2RequiresCast {
			operand2Str = fmt.Sprintf("%s(%s)", expressionGoType, operand2Str)
		}
	}

	return fmt.Sprintf("%s %s %s",
		operand1Str,
		operator,
		operand2Str)
}

func twoOperatorToGoString(scope ast.Scope, expression *ast.Expression) string {
	operator := ""
	switch expression.Type {
	case parser.ZserioLexerLSHIFT:
		operator = "<<"
	case parser.ZserioParserRSHIFT:
		operator = ">>"
	case parser.ZserioParserAND:
		operator = "&"
	case parser.ZserioParserOR:
		operator = "|"
	case parser.ZserioParserXOR:
		operator = "^"
	case parser.ZserioParserLOGICAL_AND:
		operator = "&&"
	case parser.ZserioParserLOGICAL_OR:
		operator = "||"
	}

	return fmt.Sprintf("%s %s %s",
		ExpressionToGoString(scope, expression.Operand1),
		operator,
		ExpressionToGoString(scope, expression.Operand2))
}

func ternaryExpressionToGoString(scope ast.Scope, expression *ast.Expression) string {
	op1 := ExpressionToGoString(scope, expression.Operand1)
	op2 := ExpressionToGoString(scope, expression.Operand2)
	op3 := ExpressionToGoString(scope, expression.Operand3)
	if expression.ResultType == ast.ExpressionTypeInteger {
		// in case the expression is a subtype of int (such as uint16), cast to
		// integer to avoid type cast errors
		typeCast := "int"
		op3 = fmt.Sprintf("%s(%s)", typeCast, op3)
		op2 = fmt.Sprintf("%s(%s)", typeCast, op2)
	}
	// Ternary expressions are not supported by Go. As a workaround,
	// generate the ternary expression as an if/else
	return fmt.Sprintf("%s\nif %s {\nretVal = %s\n}\n",
		op3,
		op1,
		op2,
	)
}

func lenOperatorToGoString(scope ast.Scope, expression *ast.Expression) string {
	return fmt.Sprintf("len(%s)", ExpressionToGoString(scope, expression.Operand1))
}

func valueOfOperatorToGoString(scope ast.Scope, expression *ast.Expression) string {
	return ExpressionToGoString(scope, expression.Operand1)
}

func ExpressionToGoString(scope ast.Scope, expression *ast.Expression) string {
	if expression.FullyResolved {
		switch expression.ResultType {
		case ast.ExpressionTypeInteger:
			return fmt.Sprintf("%v", expression.ResultIntValue)
		case ast.ExpressionTypeBool:
			return fmt.Sprintf("%t", expression.ResultBoolValue)
		case ast.ExpressionTypeString:
			return fmt.Sprintf("%q", expression.ResultStringValue)
		case ast.ExpressionTypeFloat:
			return fmt.Sprintf("%f", expression.ResultFloatValue)
		default:
			return "UNSUPPORTED_TYPE"
		}
	}
	switch expression.Type {
	case parser.ZserioParserLBRACKET:
		return bracketedExpressionToGoString(scope, expression)
	case parser.ZserioParserLPAREN:
		return parenthesizedExpressionToGoString(scope, expression)
	case parser.ZserioParserRPAREN:
		return functionCallExpressionToString(scope, expression)
	case parser.ZserioParserDOT:
		return dotOperatorToGoString(scope, expression)
	case parser.ZserioParserLENGTHOF:
		return lenOperatorToGoString(scope, expression)
	case parser.ZserioParserVALUEOF:
		return valueOfOperatorToGoString(scope, expression)
	case parser.ZserioParserNUMBITS:
		return numBitsOperatorToGoString(scope, expression)
	case parser.ZserioParserPLUS:
		return twoOperatorEqualTypesToGoString(scope, expression)
	case parser.ZserioParserMINUS:
		return twoOperatorEqualTypesToGoString(scope, expression)
	case parser.ZserioParserMULTIPLY:
		return twoOperatorEqualTypesToGoString(scope, expression)
	case parser.ZserioParserDIVIDE:
		return twoOperatorEqualTypesToGoString(scope, expression)
	case parser.ZserioParserBANG: // the ! (negation operator)
		return fmt.Sprintf("!%s", ExpressionToGoString(scope, expression.Operand1))
	case parser.ZserioParserMODULO:
		return twoOperatorEqualTypesToGoString(scope, expression)
	case parser.ZserioLexerLSHIFT:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserRSHIFT:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserAND:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserOR:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserXOR:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserLOGICAL_AND:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserLOGICAL_OR:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserLT:
		return twoOperatorEqualTypesToGoString(scope, expression)
	case parser.ZserioParserLE:
		return twoOperatorEqualTypesToGoString(scope, expression)
	case parser.ZserioParserGT:
		return twoOperatorEqualTypesToGoString(scope, expression)
	case parser.ZserioParserGE:
		return twoOperatorEqualTypesToGoString(scope, expression)
	case parser.ZserioParserEQ:
		return twoOperatorEqualTypesToGoString(scope, expression)
	case parser.ZserioParserNE:
		return twoOperatorEqualTypesToGoString(scope, expression)
	case parser.ZserioParserINDEX:
		// hard-code the name of the index variable used in the template
		return "index"
	case parser.ZserioParserID:
		return IdentifierToGoString(scope, expression)
	case parser.ZserioParserQUESTIONMARK:
		return ternaryExpressionToGoString(scope, expression)
	default:
		return "UNSUPPORTED"
	}
}
