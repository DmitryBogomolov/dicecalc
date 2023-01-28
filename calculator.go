package dicecalc

import "fmt"

func CalculateProbabilities(params DiceRollParameters) (*Probabilities, error) {
	if err := validateParameters(params); err != nil {
		return nil, err
	}
	min, max := getValueRange(params)
	totalCount := getVariantsCount(params)
	factorials := makeFactorials(params.DiceCount)
	count := max - min + 1
	values := make([]float64, count)
	checkCount := 0
	for i := 0; i < count; i++ {
		rolls := collectAllRolls(i+min, params)
		valueCount := calculateValueSlots(rolls, factorials)
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

func calculateValueSlots(rolls []*_DiceRoll, factorials *_Factorials) int {
	count := 0
	for _, roll := range rolls {
		k := calculateRollCount(roll, factorials)
		count += k
	}
	return count
}

func calculateRollCount(roll *_DiceRoll, factorials *_Factorials) int {
	n := len(roll.dices)
	counts := make(map[byte]int)
	for _, dice := range roll.dices {
		counts[dice]++
	}
	ret := factorials.get(n)
	for _, c := range counts {
		ret /= factorials.get(c)
	}
	return ret
}
