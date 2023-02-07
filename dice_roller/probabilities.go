package dice_roller

import "fmt"

type Probabilities struct {
	min   int
	max   int
	total int
	items []int
}

func NewProbabilities(minValue, maxValue, totalVariants int, valuesVariants []int) (*Probabilities, error) {
	if minValue > maxValue {
		return nil, fmt.Errorf("bad min value %d or max value %d", minValue, maxValue)
	}
	if len(valuesVariants) != maxValue-minValue+1 {
		return nil, fmt.Errorf("bad variants %v (!= %d)", valuesVariants, maxValue-minValue+1)
	}
	check := 0
	for _, valueVariants := range valuesVariants {
		check += valueVariants
	}
	if check != totalVariants {
		return nil, fmt.Errorf("bad total %d (!= %d)", totalVariants, check)
	}
	return &Probabilities{
		min:   minValue,
		max:   maxValue,
		total: totalVariants,
		items: valuesVariants,
	}, nil
}

func (target *Probabilities) MinValue() int {
	return target.min
}

func (target *Probabilities) MaxValue() int {
	return target.max
}

func (target *Probabilities) TotalCount() int {
	return target.total
}

func (target *Probabilities) ValuesCount() int {
	return len(target.items)
}

func (target *Probabilities) ValueCount(value int) int {
	if target.min <= value && value <= target.max {
		return target.items[value-target.min]
	}
	return 0
}

func (target *Probabilities) ValueProbability(value int) float64 {
	return float64(target.ValueCount(value)) / float64(target.total)
}
