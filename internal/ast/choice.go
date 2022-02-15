package ast

import (
	"errors"

	"github.com/woven-planet/go-zserio/internal/parser"
)

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
	SelectorExpression *Expression
	Cases              []*ChoiceCase
	DefaultCase        *ChoiceCase
	Functions          []*Function
}

func (choice *Choice) BuildScope(p *Package) error {
	// create a compound scope
	compoundScope := &CompoundScope{
		Symbol: choice,
		Scope:  make(map[string]any),
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

// fullySpecifyChoiceConditions fully specifies a choice condition in case the
// choice case is given as a text, but not a fully determined symbol.
func (choice *Choice) fullySpecifyChoiceConditions() error {
	for _, choiceCase := range choice.Cases {
		for _, condition := range choiceCase.Conditions {
			// in case the choice case is just a text string, replace that
			//  by a dot expression with the choice selector type, so that the
			// symbol is unique. If this is not done, the text value might be
			// ambiguous, for example if the value  is defined in more than one enum.
			// e.g. VALUE -> EnumType.VALUE
			if condition.Condition.Type == parser.ZserioParserID {
				// convert the selector from the parameter type to the underlying
				// actual type (e.g. enum, bitfield)
				parameterType, ok := choice.SelectorExpression.ResultSymbol.Symbol.(*Parameter)
				if !ok {
					return errors.New("choice selector type is not a parameter")
				}
				dotExpression := &Expression{
					Type: parser.ZserioParserDOT,
					Operand1: &Expression{
						Type: parser.ZserioParserID,
						Text: parameterType.Type.Name,
					},
					Operand2: condition.Condition,
				}
				condition.Condition = dotExpression
			}
		}
	}
	return nil
}
func (choice *Choice) Evaluate(p *Package) error {
	// Ignore templates
	if len(choice.TemplateParameters) > 0 {
		return nil
	}
	p.LocalSymbols.CurrentCompoundScope = &choice.Name

	// The selector expression needs to be evaluated first, because the choice
	// cases depend on the symbol refered to by the selector expression
	if err := choice.SelectorExpression.Evaluate(p); err != nil {
		return err
	}

	// Fully specify the choice condition symbols
	if err := choice.fullySpecifyChoiceConditions(); err != nil {
		return err
	}

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

	if choice.DefaultCase == nil || choice.DefaultCase.Field == nil {
		return nil
	}
	return choice.DefaultCase.Field.Evaluate(p)
}
