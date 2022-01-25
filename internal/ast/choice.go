package ast

type ChoiceCaseExpression struct {
	Condition *Expression
	Comment   string
}

type ChoiceCase struct {
	Conditions []*ChoiceCaseExpression
	Field      *Field
	Comment    string
}

type Choice struct {
	Name               string
	Comment            string
	TemplateParameters []string
	TypeParameters     []*Parameter
	Expression         *Expression
	Cases              []*ChoiceCase
	DefaultCase        *ChoiceCase
	Functions          []*Function
}

func (choice *Choice) BuildScope(p *Package) error {
	// create a compound scope
	compoundScope := &CompoundScope{
		Symbol: choice,
		Scope:  make(map[string]interface{}),
	}

	for _, param := range choice.TypeParameters {
		compoundScope.Scope[param.Name] = param
	}
	for _, choiceCase := range choice.Cases {
		if choiceCase.Field != nil {
			compoundScope.Scope[choiceCase.Field.Name] = choiceCase.Field
		}
	}
	p.LocalSymbols.CompoundScopes[choice.Name] = compoundScope
	return nil
}

func (choice *Choice) Evaluate(p *Package) error {
	// Ignore templates
	if len(choice.TemplateParameters) > 0 {
		return nil
	}
	p.LocalSymbols.CurrentCompoundScope = &choice.Name

	for _, param := range choice.TypeParameters {
		if err := param.Type.Evaluate(p); err != nil {
			return err
		}
	}
	for _, choiceCase := range choice.Cases {
		if choiceCase.Field != nil {
			if err := choiceCase.Field.Evaluate(p); err != nil {
				return err
			}
		}
	}
	for _, choiceCase := range choice.Cases {
		for _, condition := range choiceCase.Conditions {
			if err := condition.Condition.Evaluate(p); err != nil {
				return err
			}
		}
	}
	return choice.Expression.Evaluate(p)
}
