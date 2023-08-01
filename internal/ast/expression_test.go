package ast_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woven-planet/go-zserio/internal/ast"
	"github.com/woven-planet/go-zserio/internal/parser"
)

func TestStringLiteralExpressions(t *testing.T) {
	tests := map[string]struct {
		testExpression *ast.Expression
		expectedValue  string
		expectedError  error
	}{
		"simple-string": {
			testExpression: &ast.Expression{
				Type: parser.ZserioParserSTRING_LITERAL,
				Text: "\"TestString\"",
			},
			expectedValue: "TestString",
		},
		"empty-string": {
			testExpression: &ast.Expression{
				Type: parser.ZserioParserSTRING_LITERAL,
				Text: "\"\"",
			},
			expectedValue: "",
		},
		"invalid-string": {
			testExpression: &ast.Expression{
				Type: parser.ZserioParserSTRING_LITERAL,
				Text: "\"",
			},
			expectedError: errors.New("evaluate \"\\\"\": string literal should contain quotes"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			scope := &ast.Package{}
			err := test.testExpression.Evaluate(scope)
			if test.expectedError != nil {
				require.NotNil(t, err)
				assert.EqualError(t, err, test.expectedError.Error())
				return
			}
			assert.Equal(t, ast.ExpressionTypeString, test.testExpression.ResultType)
			assert.Equal(t, ast.EvaluationStateComplete, test.testExpression.EvaluationState)
			assert.Equal(t, test.expectedValue, test.testExpression.ResultStringValue)
		})
	}
}

func TestBinaryLiteralExpressions(t *testing.T) {
	tests := map[string]struct {
		testExpression *ast.Expression
		expectedValue  int64
		expectedError  error
	}{
		"simple-number": {
			testExpression: &ast.Expression{
				Type: parser.ZserioParserBINARY_LITERAL,
				Text: "111b",
			},
			expectedValue: 7,
		},
		"large-number": {
			testExpression: &ast.Expression{
				Type: parser.ZserioParserBINARY_LITERAL,
				Text: "11011110101011011011111011101111b",
			},
			expectedValue: 3735928559,
		},
		"zero": {
			testExpression: &ast.Expression{
				Type: parser.ZserioParserBINARY_LITERAL,
				Text: "0b",
			},
			expectedValue: 0,
		},
		"invalid-text": {
			testExpression: &ast.Expression{
				Type: parser.ZserioParserBINARY_LITERAL,
				Text: "10101",
			},
			expectedError: errors.New("evaluate \"10101\": binary expression is not valid"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			scope := &ast.Package{}
			err := test.testExpression.Evaluate(scope)
			if test.expectedError != nil {
				require.NotNil(t, err)
				assert.EqualError(t, err, test.expectedError.Error())
				return
			}
			assert.Equal(t, ast.ExpressionTypeInteger, test.testExpression.ResultType)
			assert.Equal(t, ast.EvaluationStateComplete, test.testExpression.EvaluationState)
			assert.Equal(t, test.expectedValue, test.testExpression.ResultIntValue)
		})
	}
}

func TestFloatArithmeticExpressions(t *testing.T) {
	tests := map[string]struct {
		testOperand1  *ast.Expression
		testOperand2  *ast.Expression
		operand       int
		expectedValue float64
	}{
		"adding-two-floats": {
			testOperand1: &ast.Expression{
				Type: parser.ZserioLexerDOUBLE_LITERAL,
				Text: "17.23",
			},
			testOperand2: &ast.Expression{
				Type: parser.ZserioLexerDOUBLE_LITERAL,
				Text: "0.5",
			},
			operand:       parser.ZserioParserPLUS,
			expectedValue: 17.73,
		},
		"subtracting-float-from-int": {
			testOperand1: &ast.Expression{
				Type: parser.ZserioLexerDECIMAL_LITERAL,
				Text: "23",
			},
			testOperand2: &ast.Expression{
				Type: parser.ZserioLexerDOUBLE_LITERAL,
				Text: "0.5",
			},
			operand:       parser.ZserioParserMINUS,
			expectedValue: 22.5,
		},
		"multiplying-float-with-int": {
			testOperand1: &ast.Expression{
				Type: parser.ZserioLexerDOUBLE_LITERAL,
				Text: "2.25",
			},
			testOperand2: &ast.Expression{
				Type: parser.ZserioLexerDECIMAL_LITERAL,
				Text: "3",
			},
			operand:       parser.ZserioParserMULTIPLY,
			expectedValue: 6.75,
		},
		"dividing-floats": {
			testOperand1: &ast.Expression{
				Type: parser.ZserioLexerDOUBLE_LITERAL,
				Text: "9.9",
			},
			testOperand2: &ast.Expression{
				Type: parser.ZserioLexerDOUBLE_LITERAL,
				Text: "3.3",
			},
			operand:       parser.ZserioParserDIVIDE,
			expectedValue: 3.0,
		},
		"do-not-divide-by-zero": {
			testOperand1: &ast.Expression{
				Type: parser.ZserioLexerDOUBLE_LITERAL,
				Text: "9.9",
			},
			testOperand2: &ast.Expression{
				Type:            parser.ZserioParserID,
				FullyResolved:   false,
				Text:            "some_identifier",
				EvaluationState: ast.EvaluationStateComplete,
				ResultType:      ast.ExpressionTypeInteger,
				ResultIntValue:  0,
			},
			operand:       parser.ZserioParserDIVIDE,
			expectedValue: 1.0,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			scope := &ast.Package{}
			expression := &ast.Expression{
				Operand1: test.testOperand1,
				Operand2: test.testOperand2,
				Type:     test.operand,
			}
			err := expression.Evaluate(scope)
			assert.Nil(t, err)
			assert.Equal(t, ast.ExpressionTypeFloat, expression.ResultType)
			assert.Equal(t, ast.EvaluationStateComplete, expression.EvaluationState)
			assert.InDelta(t, test.expectedValue, expression.ResultFloatValue, 1e-7)
		})
	}
}
