package dicecalc

type DiceSchema struct {
	Sides int
	Count int
}

type Probabilities struct {
	min    int
	max    int
	values map[int]float64
}

func (probabilities *Probabilities) Min() int {
	return probabilities.min
}

func (probabilities *Probabilities) Max() int {
	return probabilities.max
}

func (probabilities *Probabilities) Probability(value int) float64 {
	if probabilities.min <= value && value <= probabilities.max {
		return probabilities.values[value]
	}
	return 0
}

func CalculateProbabilities(schema DiceSchema) Probabilities {
	min := schema.Count
	max := schema.Count * schema.Sides
	return Probabilities{
		min: min,
		max: max,
	}
}
