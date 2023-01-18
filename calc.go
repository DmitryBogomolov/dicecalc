package dicecalc

import "math"

type DiceSchema struct {
	Sides int
	Count int
}

type Probabilities struct {
	min    int
	max    int
	values []float64
}

func (probabilities *Probabilities) Min() int {
	return probabilities.min
}

func (probabilities *Probabilities) Max() int {
	return probabilities.max
}

func (probabilities *Probabilities) Probability(value int) float64 {
	if probabilities.min <= value && value <= probabilities.max {
		return probabilities.values[value+probabilities.min]
	}
	return 0
}

const MAX_DICE_COUNT = 64
const MAX_DICE_SIDES = 32

func CalculateProbabilities(schema DiceSchema) Probabilities {
	if schema.Count < 1 || schema.Count > MAX_DICE_COUNT {
		panic("bad count") // Or err?
	}
	if schema.Sides < 1 || schema.Sides > MAX_DICE_SIDES {
		panic("bad sides") // Or err?
	}
	min := schema.Count
	max := schema.Count * schema.Sides
	count := max - min + 1
	slots := make([]int, count)
	for i := 0; i < count; i++ {
		slots[i] = calculateSlot(min+i, schema)
	}
	total := int(math.Pow(float64(schema.Sides), float64(schema.Count)))
	return Probabilities{
		min:    min,
		max:    max,
		values: calculateValues(slots, total),
	}
}

func calculateFactorials(count int) []int {
	values := make([]int, count)
	values[0] = 1
	for i := 1; i < count; i++ {
		values[i] = values[i-1] * (i + 1)
	}
	return values
}

func initSlotDices(value int, schema DiceSchema) []int {
	items := make([]int, schema.Count)
	for i := range items {
		items[i] = 1
	}
	rest := value - schema.Count
	k := schema.Count - 1
	for rest > 0 {
		portion := schema.Sides - 1
		if portion > rest {
			portion = rest
		}
		rest -= portion
		items[k] += portion
		k--
	}
	return items
}

func calculateSlot(value int, schema DiceSchema) int {
	initial := initSlotDices(value, schema)
	for k := len(initial) - 1; k > 0; k-- {
		copy := append([]int(nil), initial...)
		for copy[k] > copy[k-1] {
			sample := copy[k] - 1
			for i := k - 1; i >= 0; i-- {
			}
		}
	}
	return 0
}

func calculateValues(slots []int, totalCount int) []float64 {
	check := 0
	for _, slot := range slots {
		check += slot
	}
	if check != totalCount {
		panic("check") // Or err?
	}
	values := make([]float64, len(slots))
	for i, slot := range slots {
		values[i] = float64(slot) / float64(totalCount)
	}
	return values
}
