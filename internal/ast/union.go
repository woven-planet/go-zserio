package ast

type Union struct {
	Name               string
	Comment            string
	TemplateParameters []string
	TypeParameters     []*Parameter
	Fields             []*Field
	Functions          []*Function
}

func (u *Union) BuildScope(p *Package) error {
	// create a compound scope
	compoundScope := &CompoundScope{
		Symbol: u,
		Scope:  make(map[string]interface{}),
	}
	for _, param := range u.TypeParameters {
		compoundScope.Scope[param.Name] = param
	}
	for _, field := range u.Fields {
		compoundScope.Scope[field.Name] = field
	}
	p.LocalSymbols.CompoundScopes[u.Name] = compoundScope
	return nil
}

func (u *Union) Evaluate(p *Package) error {
	// Ignore templates
	if len(u.TemplateParameters) > 0 {
		return nil
	}
	for _, param := range u.TypeParameters {
		if err := param.Type.Evaluate(p); err != nil {
			return err
		}
	}
	for _, field := range u.Fields {
		if err := field.Evaluate(p); err != nil {
			return err
		}
	}
	for _, function := range u.Functions {
		if err := function.Result.Evaluate(p); err != nil {
			return err
		}
	}
	return nil
}
