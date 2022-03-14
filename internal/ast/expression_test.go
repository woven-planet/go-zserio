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
