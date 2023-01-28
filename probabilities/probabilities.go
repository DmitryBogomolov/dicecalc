package probabilities

type Probabilities struct {
	min   int
	max   int
	total int
	items []int
}

func NewProbabilities(min, max, total int, items []int) *Probabilities {
	return &Probabilities{
		min:   min,
		max:   max,
		total: total,
		items: items,
	}
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
