package ast

// Const is a zserio constant
type Const struct {
	Name            string
	Comment         string
	Type            *TypeReference
	ValueExpression *Expression
}

// Evaluate evalutes the value expression of a const type
func (c *Const) Evaluate(scope *Package) error {
	return c.ValueExpression.Evaluate(scope)
}
