package probabilities

import (
	"fmt"
	"math"
)

type _Probabilities struct {
	minValue       int
	maxValue       int
	minProbability float64
	maxProbability float64
	total          uint64
	items          []int
}

type Probabilities interface {
	MinValue() int
	MaxValue() int
	MinProbability() float64
	MaxProbability() float64
	VariantsCount() uint64
	ValuesCount() int
	ValueVariants(value int) int
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
	minVariant := uint64(math.MaxUint64)
	maxVariant := uint64(0)
	for _, valueVariants := range valuesVariants {
		variant := uint64(valueVariants)
		check += variant
		if variant < minVariant {
			minVariant = variant
		}
		if variant > maxVariant {
			maxVariant = variant
		}
	}
	if check != totalVariants {
		return nil, fmt.Errorf("bad total %d - should be %d)", check, totalVariants)
	}
	return &_Probabilities{
		minValue:       minValue,
		maxValue:       maxValue,
		minProbability: float64(minVariant) / float64(totalVariants),
		maxProbability: float64(maxVariant) / float64(totalVariants),
		total:          totalVariants,
		items:          valuesVariants,
	}, nil
}

func (target *_Probabilities) MinValue() int {
	return target.minValue
}

func (target *_Probabilities) MaxValue() int {
	return target.maxValue
}

func (target *_Probabilities) MinProbability() float64 {
	return target.minProbability
}

func (target *_Probabilities) MaxProbability() float64 {
	return target.maxProbability
}

func (target *_Probabilities) VariantsCount() uint64 {
	return target.total
}

func (target *_Probabilities) ValuesCount() int {
	return len(target.items)
}

func (target *_Probabilities) ValueVariants(value int) int {
	if target.minValue <= value && value <= target.maxValue {
		return target.items[value-target.minValue]
	}
	return 0
}

func (target *_Probabilities) ValueProbability(value int) float64 {
	return float64(target.ValueVariants(value)) / float64(target.total)
}
