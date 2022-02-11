package ast

import (
	"fmt"
	"errors"
	"math/bits"
	"strconv"
	"strings"

	"github.com/woven-planet/go-zserio/internal/parser"
)

// ExpressionType is the type of an expression
type ExpressionType string

const (
	ExpressionTypeInteger  ExpressionType = "int"
	ExpressionTypeFloat    ExpressionType = "float"
	ExpressionTypeString   ExpressionType = "string"
	ExpressionTypeBool     ExpressionType = "bool"
	ExpressionTypeBitmask  ExpressionType = "bitmask"
	ExpressionTypeEnum     ExpressionType = "enum"
	ExpressionTypeCompound ExpressionType = "compound"
)

// EvaluationState is the type of the expression evaluation state
type EvaluationState int

const (
	EvaluationStateNotStarted EvaluationState = iota
	EvaluationStateInProgress
	EvaluationStateComplete
)

// Expression satisfies IExpressionContext
type Expression struct {
	// Type is the type of an expression, as defined in the ZserioParser.
	Type int

	// Text is the text of the expression.
	Text string

	// Operand1 is the first operand of the expression (if used).
	Operand1 *Expression

	// Operand2 is the second operand of the expression (if used).
	Operand2 *Expression

	// Operand3 is the third operand of the expression (if used).
	Operand3 *Expression

	// ResultType is the type of the evaluated expression, e.g. ExpressionTypeInteger.
	ResultType ExpressionType

	// EvaluationState is the evaluation state of the expression.
	EvaluationState EvaluationState

	// ResultIntValue is the value of the expression if ResultType is
	// ExpressionTypeInteger or ExpressionTypeBitmask.
	ResultIntValue int64

	// ResultFloatValue is the value of the expression if ResutlType is
	// ExpressionTypeFloat.
	ResultFloatValue float64

	// ResultStringValue is the value of the expression if ResultType is
	// ExpressionTypeString.
	ResultStringValue string

	// ResultBoolValue is the value of the expression if ResultType is
	// ExpressionTypeBool
	ResultBoolValue bool

	// ResultSymbol points to the symbol if ResultType is ExpressionTypeEnum or
	// ExpressionTypeStruct.
	ResultSymbol *SymbolReference

	// FullyResolved specifies if the expression is fully resolved (no variables)
	FullyResolved bool
}

func copyExpressionResult(src, dst *Expression) {
	dst.ResultBoolValue = src.ResultBoolValue
	dst.ResultIntValue = src.ResultIntValue
	dst.ResultFloatValue = src.ResultFloatValue
	dst.ResultStringValue = src.ResultStringValue
	dst.ResultSymbol = src.ResultSymbol
	dst.ResultType = src.ResultType
}

func evaluateExpressionType(typeRef *TypeReference, scope *Package) (ExpressionType, error) {
	if !typeRef.IsBuiltin && typeRef.Package == "" {
		return "", errors.New("type is not fully resolved")
	}
	if typeRef.IsBuiltin {
		switch typeRef.Name {
		case "bool":
			return ExpressionTypeBool, nil
		case "string":
			return ExpressionTypeString, nil
		case "float16":
			return ExpressionTypeFloat, nil
		case "float32":
			return ExpressionTypeFloat, nil
		case "float64":
			return ExpressionTypeFloat, nil
		default:
			return ExpressionTypeInteger, nil
		}
	}
	symbolScope, err := scope.GetImportedScope(typeRef.Package)
	if err != nil {
		return "", err
	}
	symbol, err := symbolScope.GetSymbol(typeRef.Name)
	if err != nil {
		return "", err
	}

	return evaluateSymbolType(symbol.Symbol, scope)
}

func evaluateSymbolType(symbol interface{}, scope *Package) (ExpressionType, error) {

	switch n := symbol.(type) {
	case *Enum:
		return ExpressionTypeEnum, nil
	case *EnumItem:
		return ExpressionTypeInteger, nil
	case *Union:
		return ExpressionTypeCompound, nil
	case *Struct:
		return ExpressionTypeCompound, nil
	case *Field:
		return evaluateExpressionType(n.Type, scope)
	case *Parameter:
		return evaluateExpressionType(n.Type, scope)
	case *Subtype:
		return evaluateExpressionType(n.Type, scope)
	case *Function:
		return n.Result.ResultType, nil
	case *BitmaskType:
		return ExpressionTypeBitmask, nil
	default:
		return "", errors.New("unable to evaluate the expression type")
	}
}

// evaluateIdentifier evaluates an identifier expression.
func (expr *Expression) evaluateIdentifier(scope *Package) error {
	symbol, err := scope.GetSymbol(expr.Text)
	if err != nil {
		return errors.New("symbol not found")
	}

	// Some symbols, such as consts, can be directly resolved.
	if constSymbol, ok := symbol.Symbol.(*Const); ok {
		constSymbol.ValueExpression.Evaluate(scope)
		copyExpressionResult(constSymbol.ValueExpression, expr)
		return nil
	}

	if enumItem, ok := symbol.Symbol.(*EnumItem); ok {
		enumItem.Evaluate(scope)
		copyExpressionResult(enumItem.Expression, expr)
		return nil
	}

	// All subsequent symbols cannot be fully resolved.
	expr.FullyResolved = false
	expr.ResultSymbol = symbol
	expr.ResultType, err = evaluateSymbolType(symbol.Symbol, scope)
	return err
}

// evaluateNumBitsOperator evaluates a bit counter operator.
func (expr *Expression) evaluateNumBitsOperator() error {
	if expr.Operand1 == nil {
		return errors.New("numbits operator needs one operand")
	}
	if expr.Operand1.ResultType != ExpressionTypeInteger {
		return errors.New("numbits operator needs an integer type")
	}
	expr.ResultType = ExpressionTypeInteger
	expr.ResultIntValue = int64(bits.Len64(uint64(expr.Operand1.ResultIntValue)))
	return nil
}

func (expr *Expression) evaluateCompoundDotExpression(scope *Package) error {
	var typeReference *TypeReference
	if t, ok := expr.Operand1.ResultSymbol.Symbol.(*TypeReference); ok {
		typeReference = t
	} else if f, ok := expr.Operand1.ResultSymbol.Symbol.(*Field); ok {
		typeReference = f.Type
	} else if p, ok := expr.Operand1.ResultSymbol.Symbol.(*Parameter); ok {
		typeReference = p.Type
	} else {
		return errors.New("unknown compound dot expression")
	}

	// look up through all levels of indirections (subtypes)
	var compoundSymbol *SymbolReference
	newScope := scope
	var err error
	for {
		compoundSymbol, err = scope.GetSymbol(typeReference.Name)
		if err != nil {
			return err
		}
		newScope, err = newScope.GetImportedScope(compoundSymbol.Package)
		if err != nil {
			return err
		}
		if subtype, ok := compoundSymbol.Symbol.(*Subtype); ok {
			typeReference = subtype.Type
			// TODO comment
			newScope, err = newScope.GetImportedScope(typeReference.Package)
			if err != nil {
				return err
			}
			continue
		}
		break
	}

	if _, ok := compoundSymbol.Symbol.(*Struct); ok {
		symbol, err := scope.GetCompoundType(compoundSymbol.Name, expr.Operand2.Text)
		if err != nil {
			return err
		}
		expr.ResultSymbol = symbol
		expr.ResultType, err = evaluateSymbolType(symbol.Symbol, newScope)
		return err
	}
	return errors.New("compound type is not supported")
}

func (expr *Expression) evaluateFunctionCallExpression(scope *Package) error {
	function, ok := expr.Operand1.ResultSymbol.Symbol.(*Function)
	if !ok {
		return errors.New("function expression expected")
	}

	// find the compound type that defines the function
	compoundTypeRef := &TypeReference{
		Name:      expr.Operand1.ResultSymbol.CompoundName,
		Package:   expr.Operand1.ResultSymbol.Package,
		IsBuiltin: false,
	}
	typeRef, err := scope.GetZserioNativeType(compoundTypeRef)
	if err != nil {
		return err
	}

	// the function might not have been evaluated, hence we must evaluate it
	// now. However, to evaluate the function, a different scope needs to be
	// used.
	functionEvaluationScope, err := scope.GetImportedScope(typeRef.Type.Package)
	if err != nil {
		return err
	}
	functionEvaluationScope.LocalSymbols.CurrentCompoundScope = &typeRef.Type.Name
	err = function.Result.Evaluate(functionEvaluationScope)
	if err != nil {
		return err
	}
	copyExpressionResult(function.Result, expr)
	// copy the result symbol, in case the type of the symbol needs to be
	// evaluated.
	expr.ResultSymbol = expr.Operand1.ResultSymbol
	return nil
}

func getArrayField(expr *Expression) (*Field, error) {
	field, ok := expr.ResultSymbol.Symbol.(*Field)
	if !ok {
		return nil, errors.New("unexpected array type (field expected)")
	}

	if field.Array == nil {
		return nil, errors.New("field is not an array")
	}
	return field, nil
}

func (expr *Expression) evaluateArrayElement(scope *Package) error {
	// the object must be an array
	field, err := getArrayField(expr.Operand1)
	if err != nil {
		return err
	}
	if expr.Operand2.ResultType != ExpressionTypeInteger {
		return errors.New("array index expression must be an integger")
	}
	if expr.ResultType, err = evaluateExpressionType(field.Type, scope); err != nil {
		return err
	}
	return nil
}

func (expr *Expression) evaluateLengthOfOperator(scope *Package) error {
	_, err := getArrayField(expr.Operand1)
	if err != nil {
		return err
	}
	expr.ResultType = ExpressionTypeInteger
	return nil
}

// evaluateDotExpression evaluates a dot expression, such as:
// <enum>.<value>.
func (expr *Expression) evaluateDotExpression(scope *Package) error {
	if expr.Operand1 == nil || expr.Operand2 == nil {
		return errors.New("dot expression needs two operands")
	}

	if expr.Operand1.ResultType == ExpressionTypeEnum {
		if enum, ok := expr.Operand1.ResultSymbol.Symbol.(*Enum); ok {
			for _, enumItem := range enum.Items {
				if enumItem.Name == expr.Operand2.Text {
					if enumItem.EvaluationState != EvaluationStateComplete {
						if err := enum.Evaluate(scope); err != nil {
							return fmt.Errorf("recursive evaluation of %q: %w", enum.Name, err)
						}
					}
					copyExpressionResult(enumItem.Expression, expr)
					return nil
				}
			}
		}
	} else if expr.Operand1.ResultType == ExpressionTypeBitmask {
		if bitmask, ok := expr.Operand1.ResultSymbol.Symbol.(*BitmaskType); ok {
			for _, bitmaskValue := range bitmask.Values {
				if bitmaskValue.Name == expr.Operand2.Text {
					copyExpressionResult(bitmaskValue.Expression, expr)
					return nil
				}
			}
		}
	} else if expr.Operand1.ResultType == ExpressionTypeCompound {
		return expr.evaluateCompoundDotExpression(scope)
	} else {
		return errors.New("dot expression must use a valid symbol")
	}
	return errors.New("failed to evaluate dot expression")
}

// evaluateParenthesizedExpression evaluates an expression inside a parentheses.
func (expr *Expression) evaluateParenthesizedExpression() error {
	if expr.Operand1 == nil {
		return errors.New("parenthesis needs an operand")
	}
	expr.ResultType = expr.Operand1.ResultType
	expr.ResultIntValue = expr.Operand1.ResultIntValue
	expr.ResultStringValue = expr.Operand1.ResultStringValue
	expr.ResultBoolValue = expr.Operand1.ResultBoolValue
	return nil
}

// evaluateUnaryArithmeticExpression calculates an unary arithmetic expression,
// such as -5, or +6.35. Calling it with a different operand type than int
// or float will trigger an error.
func (expr *Expression) evaluateUnaryArithmeticExpression() error {
	if expr.Type != parser.ZserioParserPLUS && expr.Type != parser.ZserioParserMINUS {
		return errors.New("invalid unary arithmetic expression")
	}
	if expr.Operand1.ResultType != ExpressionTypeInteger &&
		expr.Operand1.ResultType != ExpressionTypeFloat {
		return errors.New("unary arithmetic expression must be float or int")
	}
	expr.ResultType = expr.Operand1.ResultType
	if expr.Operand1.ResultType == ExpressionTypeInteger {
		expr.ResultIntValue = expr.Operand1.ResultIntValue
		if expr.Type == parser.ZserioParserMINUS {
			expr.ResultIntValue = -expr.Operand1.ResultIntValue
		}
	}
	return nil
}

// evaluateArithmeticExpression evaluates an arithmetic expression: plus, minus,
// multiply, divide and modulo.
// The accepted data types are string, int, and float, with one or two operands.
// invalid combinations of operands (e.g. multiply two strings) will return an
// error.
func (expr *Expression) evaluateArithmeticExpression() error {
	if expr.Operand1 == nil {
		return errors.New("arithmetic operations need at least one operand")
	}
	// single arithemtic expressions
	if expr.Operand2 == nil {
		return expr.evaluateUnaryArithmeticExpression()
	}

	// Check is string addition is wanted
	if expr.Operand1.ResultType == ExpressionTypeString &&
		expr.Operand2.ResultType == ExpressionTypeString {
		if expr.Type == parser.ZserioParserPLUS {
			expr.ResultType = ExpressionTypeString
			expr.ResultStringValue = expr.Operand1.ResultStringValue + expr.Operand2.ResultStringValue
			return nil
		}
		return errors.New("invalid operation on string operands")
	}

	// currently, only integer and string arithmetics are supported
	if expr.Operand1.ResultType != ExpressionTypeInteger || expr.Operand2.ResultType != ExpressionTypeInteger {
		return errors.New("arithmetic expressions for types other than int or string not implemented")
	}

	expr.ResultType = ExpressionTypeInteger
	switch expr.Type {
	case parser.ZserioParserPLUS:
		expr.ResultIntValue = expr.Operand1.ResultIntValue + expr.Operand2.ResultIntValue
	case parser.ZserioParserMINUS:
		expr.ResultIntValue = expr.Operand1.ResultIntValue - expr.Operand2.ResultIntValue
	case parser.ZserioParserMULTIPLY:
		expr.ResultIntValue = expr.Operand1.ResultIntValue * expr.Operand2.ResultIntValue
	case parser.ZserioParserDIVIDE:
		expr.ResultIntValue = expr.Operand1.ResultIntValue / expr.Operand2.ResultIntValue
	case parser.ZserioParserMODULO:
		expr.ResultIntValue = expr.Operand1.ResultIntValue % expr.Operand2.ResultIntValue
	default:
		return errors.New("unexpected operation in arithmetic expression")
	}
	return nil
}

func (expr *Expression) evaluateLogicalNegation() error {
	if expr.Operand1 == nil {
		return errors.New("negation needs exactly one operand")
	}
	if expr.Operand1.ResultType != ExpressionTypeBool {
		return errors.New("boolean operand expected")
	}
	expr.ResultType = expr.Operand1.ResultType
	expr.ResultBoolValue = !expr.Operand1.ResultBoolValue
	return nil
}

func (expr *Expression) evaluateLogicalExpression() error {
	if expr.Operand1 == nil || expr.Operand2 == nil {
		return errors.New("bitwise expressions always need to operands")
	}

	if expr.Operand1.ResultType != ExpressionTypeBool ||
		expr.Operand2.ResultType != ExpressionTypeBool {
		return errors.New("integer or bitmask types expected for logical expressions")
	}

	expr.ResultType = expr.Operand1.ResultType
	switch expr.Type {
	case parser.ZserioParserLOGICAL_AND:
		expr.ResultBoolValue = expr.Operand1.ResultBoolValue && expr.Operand2.ResultBoolValue
	case parser.ZserioParserLOGICAL_OR:
		expr.ResultBoolValue = expr.Operand1.ResultBoolValue || expr.Operand2.ResultBoolValue
	default:
		return errors.New("unexpected logical operand")
	}
	return nil
}

func (expr *Expression) evaluateComparisonExpression() error {
	if expr.Operand1.ResultType != expr.Operand2.ResultType &&
		!(expr.Operand1.ResultType == ExpressionTypeEnum &&
			expr.Operand2.ResultType == ExpressionTypeInteger) &&
		!(expr.Operand2.ResultType == ExpressionTypeEnum &&
			expr.Operand1.ResultType == ExpressionTypeInteger) &&
		!(expr.Operand2.ResultType == ExpressionTypeBitmask &&
			expr.Operand1.ResultType == ExpressionTypeInteger) &&
		!(expr.Operand1.ResultType == ExpressionTypeBitmask &&
			expr.Operand2.ResultType == ExpressionTypeInteger) {
		return errors.New("comparison expression must have operands of the same type")
	}
	expr.ResultType = ExpressionTypeBool
	if expr.Operand1.ResultType == ExpressionTypeInteger ||
		expr.Operand1.ResultType == ExpressionTypeBitmask {
		switch expr.Type {
		case parser.ZserioParserLT:
			expr.ResultBoolValue = expr.Operand1.ResultIntValue < expr.Operand2.ResultIntValue
		case parser.ZserioParserLE:
			expr.ResultBoolValue = expr.Operand1.ResultIntValue <= expr.Operand2.ResultIntValue
		case parser.ZserioParserGT:
			expr.ResultBoolValue = expr.Operand1.ResultIntValue > expr.Operand2.ResultIntValue
		case parser.ZserioParserGE:
			expr.ResultBoolValue = expr.Operand1.ResultIntValue >= expr.Operand2.ResultIntValue
		case parser.ZserioParserEQ:
			expr.ResultBoolValue = expr.Operand1.ResultIntValue == expr.Operand2.ResultIntValue
		case parser.ZserioParserNE:
			expr.ResultBoolValue = expr.Operand1.ResultIntValue != expr.Operand2.ResultIntValue
		default:
			return errors.New("unexpected operand for integer expression")
		}
	} else if expr.Operand1.ResultType == ExpressionTypeFloat {
		switch expr.Type {
		case parser.ZserioParserLT:
			expr.ResultBoolValue = expr.Operand1.ResultFloatValue < expr.Operand2.ResultFloatValue
		case parser.ZserioParserLE:
			expr.ResultBoolValue = expr.Operand1.ResultFloatValue <= expr.Operand2.ResultFloatValue
		case parser.ZserioParserGT:
			expr.ResultBoolValue = expr.Operand1.ResultFloatValue > expr.Operand2.ResultFloatValue
		case parser.ZserioParserGE:
			expr.ResultBoolValue = expr.Operand1.ResultFloatValue >= expr.Operand2.ResultFloatValue
		default:
			return errors.New("equal/notequal operator not supported for float types")
		}
	} else if expr.Operand1.ResultType == ExpressionTypeString {
		switch expr.Type {
		case parser.ZserioParserEQ:
			expr.ResultBoolValue = expr.Operand1.ResultStringValue == expr.Operand2.ResultStringValue
		case parser.ZserioParserNE:
			expr.ResultBoolValue = expr.Operand1.ResultStringValue != expr.Operand2.ResultStringValue
		default:
			return errors.New("string types does not support greater / smaller comparisons")
		}
	}
	return nil
}

func (expr *Expression) evaluateTernaryExpression() error {
	if expr.Operand1.ResultType != ExpressionTypeBool {
		return errors.New("ternary expression operand 1 must be a boolean expression")
	}

	if expr.Operand1.ResultBoolValue {
		expr.ResultType = expr.Operand2.ResultType
		expr.ResultIntValue = expr.Operand2.ResultIntValue
		expr.ResultBoolValue = expr.Operand2.ResultBoolValue
		expr.ResultStringValue = expr.Operand2.ResultStringValue
	} else {
		expr.ResultType = expr.Operand3.ResultType
		expr.ResultIntValue = expr.Operand3.ResultIntValue
		expr.ResultBoolValue = expr.Operand3.ResultBoolValue
		expr.ResultStringValue = expr.Operand3.ResultStringValue
	}
	return nil
}

func (expr *Expression) evaluateBitwiseNegation() error {
	if expr.Operand1 == nil {
		return errors.New("negation needs exactly one operand")
	}
	if expr.Operand1.ResultType != ExpressionTypeInteger &&
		expr.Operand1.ResultType != ExpressionTypeBitmask {
		return errors.New("integer or bitmask type expected")
	}
	expr.ResultType = expr.Operand1.ResultType
	expr.ResultIntValue = ^expr.Operand1.ResultIntValue
	return nil
}

func (expr *Expression) evaluateBitwiseExpression() error {
	if expr.Operand1 == nil || expr.Operand2 == nil {
		return errors.New("bitwise expressions always need to operands")
	}

	if (expr.Operand1.ResultType != ExpressionTypeBitmask && expr.Operand1.ResultType != ExpressionTypeInteger) ||
		(expr.Operand2.ResultType != ExpressionTypeBitmask && expr.Operand2.ResultType != ExpressionTypeInteger) {
		return errors.New("integer or bitmask types expected for bitwise expressions")
	}

	expr.ResultType = expr.Operand1.ResultType

	switch expr.Type {
	case parser.ZserioLexerLSHIFT:
		expr.ResultIntValue = expr.Operand1.ResultIntValue << expr.Operand2.ResultIntValue
	case parser.ZserioParserRSHIFT:
		expr.ResultIntValue = expr.Operand1.ResultIntValue >> expr.Operand2.ResultIntValue
	case parser.ZserioParserAND:
		expr.ResultIntValue = expr.Operand1.ResultIntValue & expr.Operand2.ResultIntValue
	case parser.ZserioParserOR:
		expr.ResultIntValue = expr.Operand1.ResultIntValue | expr.Operand2.ResultIntValue
	case parser.ZserioParserXOR:
		expr.ResultIntValue = expr.Operand1.ResultIntValue ^ expr.Operand2.ResultIntValue
	default:
		return errors.New("unexpected bitwise operator")
	}
	return nil
}

// Evaluate evalues the value of the expression, by evaluating all child
// expressions. If successful, expr will have valid ExpressionType and Values
// set.
func (expr *Expression) Evaluate(scope *Package) error {
	if expr == nil {
		panic("can not evaluate a nil expression")
	}
	var err error

	if expr.EvaluationState == EvaluationStateComplete {
		return nil
	}

	if expr.EvaluationState == EvaluationStateInProgress {
		return errors.New("cyclic expression dependency found")
	}

	expr.EvaluationState = EvaluationStateInProgress

	// First, assume that the expression can be fully resolved
	expr.FullyResolved = true
	if expr.Operand1 != nil {
		err = expr.Operand1.Evaluate(scope)
		if err != nil {
			return err
		}
		if !expr.Operand1.FullyResolved {
			expr.FullyResolved = false
		}
	}
	if expr.Operand2 != nil {
		err = expr.Operand2.Evaluate(scope)
		if err != nil {
			return err
		}
		if !expr.Operand2.FullyResolved {
			expr.FullyResolved = false
		}
	}

	if expr.Operand3 != nil {
		err = expr.Operand3.Evaluate(scope)
		if err != nil {
			return err
		}
		if !expr.Operand3.FullyResolved {
			expr.FullyResolved = false
		}
	}

	switch expr.Type {
	case parser.ZserioParserLPAREN:
		err = expr.evaluateParenthesizedExpression()
	case parser.ZserioParserRPAREN:
		err = expr.evaluateFunctionCallExpression(scope)
	case parser.ZserioParserLBRACKET:
		err = expr.evaluateArrayElement(scope)
	case parser.ZserioParserDOT:
		err = expr.evaluateDotExpression(scope)
	case parser.ZserioParserLENGTHOF:
		err = expr.evaluateLengthOfOperator(scope)
		/*

			case parser.ZserioParserVALUEOF:

		*/
	case parser.ZserioParserNUMBITS:
		err = expr.evaluateNumBitsOperator()
	case parser.ZserioParserPLUS:
		err = expr.evaluateArithmeticExpression()
	case parser.ZserioParserMINUS:
		err = expr.evaluateArithmeticExpression()
	case parser.ZserioParserMULTIPLY:
		err = expr.evaluateArithmeticExpression()
	case parser.ZserioParserDIVIDE:
		err = expr.evaluateArithmeticExpression()
	case parser.ZserioParserMODULO:
		err = expr.evaluateArithmeticExpression()
	case parser.ZserioParserBANG: // the ! (negation operator)
		err = expr.evaluateLogicalNegation()
	case parser.ZserioParserTILDE:
		err = expr.evaluateBitwiseNegation()
	case parser.ZserioLexerLSHIFT:
		err = expr.evaluateBitwiseExpression()
	case parser.ZserioParserRSHIFT:
		err = expr.evaluateBitwiseExpression()
	case parser.ZserioParserAND:
		err = expr.evaluateBitwiseExpression()
	case parser.ZserioParserOR:
		err = expr.evaluateBitwiseExpression()
	case parser.ZserioParserXOR:
		err = expr.evaluateBitwiseExpression()
	case parser.ZserioParserLOGICAL_AND:
		err = expr.evaluateLogicalExpression()
	case parser.ZserioParserLOGICAL_OR:
		err = expr.evaluateLogicalExpression()
	case parser.ZserioParserLT:
		err = expr.evaluateComparisonExpression()
	case parser.ZserioParserLE:
		err = expr.evaluateComparisonExpression()
	case parser.ZserioParserGT:
		err = expr.evaluateComparisonExpression()
	case parser.ZserioParserGE:
		err = expr.evaluateComparisonExpression()
	case parser.ZserioParserEQ:
		err = expr.evaluateComparisonExpression()
	case parser.ZserioParserNE:
		err = expr.evaluateComparisonExpression()
	case parser.ZserioParserBINARY_LITERAL:
		expr.ResultType = ExpressionTypeInteger
		// remove the trailing "b"
		binaryString := expr.Text
		if !strings.HasSuffix(binaryString, "b") {
			return errors.New("binary expression is not valid")
		}
		binaryString = binaryString[:len(binaryString)-1]
		expr.ResultIntValue, err = strconv.ParseInt(binaryString, 2, 64)
	case parser.ZserioParserQUESTIONMARK:
		err = expr.evaluateTernaryExpression()
	case parser.ZserioParserOCTAL_LITERAL:
		expr.ResultType = ExpressionTypeInteger
		// use base 0 to be able to parse the "0" prefix
		expr.ResultIntValue, err = strconv.ParseInt(expr.Text, 0, 64)
	case parser.ZserioParserDECIMAL_LITERAL:
		expr.ResultType = ExpressionTypeInteger
		expr.ResultIntValue, err = strconv.ParseInt(expr.Text, 10, 64)
	case parser.ZserioParserHEXADECIMAL_LITERAL:
		expr.ResultType = ExpressionTypeInteger
		// use base 0 to be able to parse the "0x" prefix
		expr.ResultIntValue, err = strconv.ParseInt(expr.Text, 0, 64)
	case parser.ZserioParserBOOL_LITERAL:
		expr.ResultType = ExpressionTypeBool
		expr.ResultBoolValue = false
		if expr.Text == "true" {
			expr.ResultBoolValue = true
		}
	case parser.ZserioParserID:
		err = expr.evaluateIdentifier(scope)
	case parser.UnevaluatableExpressionType:
		err = nil
	default:
		err = errors.New("unknown expression type")
	}

	if err == nil {
		expr.EvaluationState = EvaluationStateComplete
	}
	return err
}
