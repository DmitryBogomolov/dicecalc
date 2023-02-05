package sum_dice

type _Factorials struct {
	values []int
}

func makeFactorials(length int) *_Factorials {
	values := make([]int, length)
	values[0] = 1
	for i := 1; i < length; i++ {
		values[i] = values[i-1] * (i + 1)
	}
	return &_Factorials{values: values}
}

func (factorials *_Factorials) get(value int) int {
	return factorials.values[value-1]
}
