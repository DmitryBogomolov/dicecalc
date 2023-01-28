package dicecalc

type Probabilities struct {
	min    int
	max    int
	values []float64
}

func (probabilities *Probabilities) MinValue() int {
	return probabilities.min
}

func (probabilities *Probabilities) MaxValue() int {
	return probabilities.max
}

func (probabilities *Probabilities) ValuesCount() int {
	return len(probabilities.values)
}

func (probabilities *Probabilities) ValueProbability(value int) float64 {
	if probabilities.min <= value && value <= probabilities.max {
		return probabilities.values[value-probabilities.min]
	}
	return 0
}
