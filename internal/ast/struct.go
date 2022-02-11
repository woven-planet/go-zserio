package ast

type Struct struct {
	Name               string
	Comment            string
	TemplateParameters []string
	TypeParameters     []*Parameter
	Fields             []*Field
	Functions          []*Function
}

func (s *Struct) BuildScope(p *Package) error {
	// create a compound scope
	compoundScope := &CompoundScope{
		Symbol: s,
		Scope:  make(map[string]any),
	}
	for _, param := range s.TypeParameters {
		compoundScope.Scope[param.Name] = param
	}
	for _, field := range s.Fields {
		compoundScope.Scope[field.Name] = field
	}
	for _, function := range s.Functions {
		compoundScope.Scope[function.Name] = function
	}
	p.LocalSymbols.CompoundScopes[s.Name] = compoundScope
	return nil
}

func (s *Struct) Evaluate(p *Package) error {
	// Ignore templates
	if len(s.TemplateParameters) > 0 {
		return nil
	}
	p.LocalSymbols.CurrentCompoundScope = &s.Name

	for _, param := range s.TypeParameters {
		if err := param.Type.Evaluate(p); err != nil {
			return err
		}
	}
	for _, field := range s.Fields {
		if err := field.Evaluate(p); err != nil {
			return err
		}
	}

	for _, function := range s.Functions {
		if err := function.Result.Evaluate(p); err != nil {
			return err
		}
	}
	return nil
}
