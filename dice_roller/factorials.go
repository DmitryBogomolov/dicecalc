package dice_roller

type Factorials func(t int) int

func CalculateFactorials(length int) Factorials {
	values := make([]int, length)
	values[0] = 1
	for i := 1; i < length; i++ {
		values[i] = values[i-1] * (i + 1)
	}
	return func(t int) int {
		return values[(t - 1)]
	}
}
