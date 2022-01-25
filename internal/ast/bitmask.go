package ast

type BitmaskValue struct {
	Name string
	*Expression
}
type BitmaskType struct {
	Name    string
	Comment string
	Type    *TypeReference
	Values  []*BitmaskValue
}

func findNextFreeBit(lastValue int64) int64 {
	highestBitPosition := 0
	for lastValue > 0 {
		lastValue = lastValue >> 1
		highestBitPosition++
	}
	return 1 << highestBitPosition
}

func (bitmask *BitmaskType) Evaluate(p *Package) error {
	for ix, value := range bitmask.Values {
		if value.Expression == nil {
			if ix == 0 {
				value.Expression = &Expression{
					EvaluationState: EvaluationStateComplete,
					ResultType:      ExpressionTypeInteger,
					ResultIntValue:  1, // bitmasks start with 1
				}
			} else {
				if bitmask.Values[ix-1].ResultType != ExpressionTypeInteger {
					return MissingEnumValue{value.Name}
				}
				value.Expression = &Expression{
					EvaluationState: EvaluationStateComplete,
					ResultType:      ExpressionTypeInteger,
					ResultIntValue:  findNextFreeBit(bitmask.Values[ix-1].ResultIntValue),
				}
			}
		} else {
			if err := value.Evaluate(p); err != nil {
				return err
			}
		}
	}
	return nil
}
