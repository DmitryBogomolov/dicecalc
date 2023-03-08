package probabilities

import "fmt"

type _Probabilities struct {
	min   int
	max   int
	total uint64
	items []int
}

type Probabilities interface {
	MinValue() int
	MaxValue() int
	TotalCount() uint64
	ValuesCount() int
	ValueCount(value int) int
	ValueProbability(value int) float64
}

func NewProbabilities(minValue int, maxValue int, totalVariants uint64, valuesVariants []int) (Probabilities, error) {
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
	if check != totalVariants {
		return nil, fmt.Errorf("bad total %d - should be %d)", check, totalVariants)
	}
	return &_Probabilities{
		min:   minValue,
		max:   maxValue,
		total: totalVariants,
		items: valuesVariants,
	}, nil
}

func (target *_Probabilities) MinValue() int {
	return target.min
}

func (target *_Probabilities) MaxValue() int {
	return target.max
}

func (target *_Probabilities) TotalCount() uint64 {
	return target.total
}

func (target *_Probabilities) ValuesCount() int {
	return len(target.items)
}

func (target *_Probabilities) ValueCount(value int) int {
	if target.min <= value && value <= target.max {
		return target.items[value-target.min]
	}
	return 0
}

func (target *_Probabilities) ValueProbability(value int) float64 {
	return float64(target.ValueCount(value)) / float64(target.total)
}
