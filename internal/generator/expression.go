package generator

import (
	"fmt"
	"strings"

	"github.com/woven-planet/go-zserio/internal/ast"
	"github.com/woven-planet/go-zserio/internal/parser"
)

func IdentifierToGoString(expression *ast.Expression) string {
	if field, ok := expression.ResultSymbol.Symbol.(*ast.Field); ok {
		return "v." + field.Name
	}
	if parameter, ok := expression.ResultSymbol.Symbol.(*ast.Parameter); ok {
		return "v." + parameter.Name
	}
	return "UNSUPPORTED"
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
	if expression.Operand1.ResultType == ast.ExpressionTypeEnum {
		return leftText + rightText
	}
	return fmt.Sprintf("%s.%s", leftText, rightText)
}

func numBitsOperatorToGoString(scope ast.Scope, expression *ast.Expression) string {
	// TODO this is a very weak way of getting the bitsize, but fixing this
	// properly needs a major refactoring in expression handling (need to find
	// out the type (signed/unsigned) of the expression result, and perhaps
	// even the maximum/minimum bit size.
	return fmt.Sprintf("ztype.TryUnsignedBitSize(uint64(%s))", ExpressionToGoString(scope, expression.Operand1))
}

func twoOperatorToGoString(scope ast.Scope, expression *ast.Expression) string {
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
	return fmt.Sprintf("%s %s %s",
		ExpressionToGoString(scope, expression.Operand1),
		operator,
		ExpressionToGoString(scope, expression.Operand2))
}

func ternaryExpressionToGoString(scope ast.Scope, expression *ast.Expression) string {
	return fmt.Sprintf("%s\nif %s {\nretVal = %s\n}\n",
		ExpressionToGoString(scope, expression.Operand3),
		ExpressionToGoString(scope, expression.Operand1),
		ExpressionToGoString(scope, expression.Operand2))
}

func lenOperatorToGoString(scope ast.Scope, expression *ast.Expression) string {

	innerExpression := ExpressionToGoString(scope, expression.Operand1)

	if f, ok := expression.Operand1.ResultSymbol.Symbol.(*ast.Field); ok {
		if f.Array != nil {
			innerExpression += ".RawArray"
		}
	}
	return fmt.Sprintf("len(%s)", innerExpression)
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
		default:
			return "UNSUPPORTED"
		}
	}
	switch expression.Type {
	case parser.ZserioParserLPAREN:
		return parenthesizedExpressionToGoString(scope, expression)
	case parser.ZserioParserRPAREN:
		return functionCallExpressionToString(scope, expression)
	case parser.ZserioParserDOT:
		return dotOperatorToGoString(scope, expression)
	case parser.ZserioParserLENGTHOF:
		return lenOperatorToGoString(scope, expression)
	case parser.ZserioParserNUMBITS:
		return numBitsOperatorToGoString(scope, expression)
	case parser.ZserioParserPLUS:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserMINUS:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserMULTIPLY:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserDIVIDE:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserBANG: // the ! (negation operator)
		return fmt.Sprintf("!%s", ExpressionToGoString(scope, expression.Operand1))
	case parser.ZserioParserMODULO:
		return twoOperatorToGoString(scope, expression)
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
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserLE:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserGT:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserGE:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserEQ:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserNE:
		return twoOperatorToGoString(scope, expression)
	case parser.ZserioParserID:
		return IdentifierToGoString(expression)
	case parser.ZserioParserQUESTIONMARK:
		return ternaryExpressionToGoString(scope, expression)
	default:
		return "UNSUPPORTED"
	}
}
