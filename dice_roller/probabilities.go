package dice_roller

import "fmt"

// TODO: Add interface. Don't forget to add implementation check.
type Probabilities struct {
	min   int
	max   int
	total uint64
	items []int
}

func NewProbabilities(minValue int, maxValue int, totalVariants uint64, valuesVariants []int) (*Probabilities, error) {
	if minValue > maxValue {
		return nil, fmt.Errorf("bad value range: %d..%d", minValue, maxValue)
	}
	if len(valuesVariants) != maxValue-minValue+1 {
		return nil, fmt.Errorf("bad variants - length should be %d", maxValue-minValue+1)
	}
	check := uint64(0)
	for _, valueVariants := range valuesVariants {
		check += uint64(valueVariants)
	}
	// TODO: Remove this argument. Just calculate.
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

func (target *Probabilities) TotalCount() uint64 {
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
