package factorials

type Factorials struct {
	values []int
}

func New(length int) *Factorials {
	values := make([]int, length)
	values[0] = 1
	for i := 1; i < length; i++ {
		values[i] = values[i-1] * (i + 1)
	}
	return &Factorials{values: values}
}

func (factorials *Factorials) Get(value int) int {
	return factorials.values[value-1]
}
