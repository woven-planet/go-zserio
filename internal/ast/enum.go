package ast

import "fmt"

type MissingEnumValue struct {
	Name string
}

func (e MissingEnumValue) Error() string {
	return fmt.Sprintf("no value set for enum item %s", e.Name)
}

type Enum struct {
	Name    string
	Comment string
	Items   []*EnumItem
	Type    *TypeReference
}

func (e *Enum) Evaluate(scope *Package) error {
	for ix, item := range e.Items {
		if item.Expression == nil {
			if ix == 0 {
				item.Expression = &Expression{
					EvaluationState: EvaluationStateComplete,
					ResultType:      ExpressionTypeInteger,
					ResultIntValue:  0,
				}
			} else {
				if e.Items[ix-1].ResultType != ExpressionTypeInteger {
					return MissingEnumValue{item.Name}
				}
				item.Expression = &Expression{
					EvaluationState: EvaluationStateComplete,
					ResultType:      ExpressionTypeInteger,
					ResultIntValue:  e.Items[ix-1].ResultIntValue + 1,
				}
			}
		} else {
			if err := item.Evaluate(scope); err != nil {
				return err
			}
		}
	}
	return nil
}

type EnumItem struct {
	Name    string
	Comment string
	*Expression
}
