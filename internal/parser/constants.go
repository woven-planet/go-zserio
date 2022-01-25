package parser

const (
	// UnevaluatableExpressionType is as an expression type which cannot be
	// evaluated by itself. An example is the right part of a dot expression:
	// <type>.<value>, where <value> cannot be evaluated without the type.
	UnevaluatableExpressionType = 0xFFFFFF
)
