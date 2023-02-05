package sum_dice_par

import (
	"math"

	"github.com/DmitryBogomolov/dicecalc/factorials"
	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

func CalculateProbabilities(params probabilities.DiceRollParameters) (*probabilities.Probabilities, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	min := params.DiceCount
	max := params.DiceCount * params.DiceSides
	totalCount := int(math.Pow(float64(params.DiceSides), float64(params.DiceCount)))
	factorials := factorials.New(params.DiceCount)
	len := max - min + 1
	values := make([]int, len)
	half := len >> 1
	for i := 0; i <= half; i++ {
		k := i
		rolls := collectAllRolls(k+min, params)
		values[k] = calculateValueSlots(rolls, factorials)
	}
	for i := half + 1; i < len; i++ {
		values[i] = values[len-1-i]
	}
	return probabilities.NewProbabilities(min, max, totalCount, values)
}

func calculateValueSlots(rolls []*_DiceRoll, factorials *factorials.Factorials) int {
	count := 0
	for _, roll := range rolls {
		k := calculateRollCount(roll, factorials)
		count += k
	}
	return count
}

func calculateRollCount(roll *_DiceRoll, factorials *factorials.Factorials) int {
	n := len(roll.dices)
	counts := make(map[byte]int)
	for _, dice := range roll.dices {
		counts[dice]++
	}
	ret := factorials.Get(n)
	for _, c := range counts {
		ret /= factorials.Get(c)
	}
	return ret
}
