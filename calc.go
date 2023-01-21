package dicecalc

import "fmt"

func CalculateProbabilities(params DiceRollParameters) (*Probabilities, error) {
	if err := validateParameters(params); err != nil {
		return nil, err
	}
	min, max := getValueRange(params)
	totalCount := getVariantsCount(params)
	// TODO: Use precalculated metrics.
	factorials := calculateFactorials(params.DiceCount)
	count := max - min + 1
	values := make([]float64, count)
	checkCount := 0
	for i := 0; i < count; i++ {
		rolls := collectAllRolls(i+min, params)
		valueCount := 0
		for _, roll := range rolls {
			k := calculateRollCount(roll, factorials)
			valueCount += k
		}
		checkCount += valueCount
		values[i] = float64(valueCount) / float64(totalCount)
	}
	if checkCount != totalCount {
		panic(fmt.Errorf("no match: expected %d, got %d", totalCount, checkCount))
	}
	return &Probabilities{
		min:    min,
		max:    max,
		values: values,
	}, nil
}

func calculateRollCount(roll *_DiceRoll, factorials []int) int {
	n := len(roll.dices)
	counts := make(map[byte]int)
	for _, dice := range roll.dices {
		counts[dice]++
	}
	ret := factorials[n-1]
	for _, c := range counts {
		ret /= factorials[c-1]
	}
	return ret
}

func calculateFactorials(length int) []int {
	values := make([]int, length)
	values[0] = 1
	for i := 1; i < length; i++ {
		values[i] = values[i-1] * (i + 1)
	}
	return values
}
