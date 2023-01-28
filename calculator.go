package dicecalc

import (
	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

func CalculateProbabilities(params probabilities.DiceRollParameters) (*probabilities.Probabilities, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	min, max := params.GetValueRange()
	totalCount := params.GetVariantsCount()
	factorials := makeFactorials(params.DiceCount)
	values := make([]int, max-min+1)
	for i := range values {
		rolls := collectAllRolls(i+min, params)
		values[i] = calculateValueSlots(rolls, factorials)
	}
	return probabilities.NewProbabilities(min, max, totalCount, values)
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
