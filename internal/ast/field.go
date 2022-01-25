package ast

type Array struct {
	IsPacked bool
	Length   *Expression
}

type Field struct {
	Name    string
	Comment string
	// Optional fields are always serialised as a boolean that indicates
	// if the value field is set. If the field is not set no value is
	// serialised.
	// http://zserio.org/doc/ZserioLanguageOverview.html#optional-members
	IsOptional     bool
	Alignment      uint8
	Type           *TypeReference
	Initializer    *Expression
	Constraint     *Expression
	OptionalClause *Expression
	// Array stored the array properties if the field is an array
	Array *Array
}

func (f *Field) Evaluate(p *Package) error {
	if err := f.Type.Evaluate(p); err != nil {
		return err
	}
	if f.Initializer != nil {
		if err := f.Initializer.Evaluate(p); err != nil {
			return err
		}
	}
	if f.OptionalClause != nil {
		if err := f.OptionalClause.Evaluate(p); err != nil {
			return err
		}
	}
	if f.Constraint != nil {
		if err := f.Constraint.Evaluate(p); err != nil {
			return err
		}
	}
	if f.Array != nil {
		if f.Array.Length != nil {
			if err := f.Array.Length.Evaluate(p); err != nil {
				return err
			}
		}
	}
	return nil
}
