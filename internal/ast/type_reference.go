package ast

type TypeReference struct {
	// IsBuiltin is set to true for standard zserio types
	IsBuiltin         bool
	Package           string
	Name              string
	Bits              int
	TemplateArguments []*TypeReference

	// TypeArguments are the type arguments if the underlying type supports
	// parameters.
	TypeArguments []*Expression

	// LengthExpression is an optional expression, specifying the length of a bit
	// mask, for example int<Expression>. If the expression evaluates to 18, the
	// resulting type will be int<18>
	LengthExpression *Expression
}

func (tr *TypeReference) Evaluate(p *Package) error {
	if tr.LengthExpression != nil {
		if err := tr.LengthExpression.Evaluate(p); err != nil {
			return err
		}
	}
	for _, typeArgument := range tr.TypeArguments {
		if err := typeArgument.Evaluate(p); err != nil {
			return err
		}
	}
	return nil
}
