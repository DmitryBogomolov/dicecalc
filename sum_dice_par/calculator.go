package sum_dice_par

import (
	"github.com/DmitryBogomolov/dicecalc/dice_roller"
)

func CalculateProbabilities(params dice_roller.DiceRollParameters) (*dice_roller.Probabilities, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	min := params.DiceCount
	max := params.DiceCount * params.DiceSides
	roller := dice_roller.NewRoller(params)
	factorials := dice_roller.CalculateFactorials(params.DiceCount)
	len := max - min + 1
	values := make([]int, len)
	half := len >> 1
	for i := 0; i <= half; i++ {
		k := i
		rolls := collectAllRolls(k+min, roller)
		values[k] = calculateValueSlots(rolls, factorials)
	}
	for i := half + 1; i < len; i++ {
		values[i] = values[(len - 1 - i)]
	}
	return dice_roller.NewProbabilities(min, max, roller.TotalRolls(), values)
}

func calculateValueSlots(rolls []dice_roller.DiceRoll, factorials dice_roller.Factorials) int {
	count := 0
	for _, roll := range rolls {
		k := calculateRollCount(roll, factorials)
		count += k
	}
	return count
}

func calculateRollCount(roll dice_roller.DiceRoll, factorials dice_roller.Factorials) int {
	n := len(roll)
	counts := make(map[int]int)
	for _, dice := range roll {
		counts[dice]++
	}
	ret := factorials(n)
	for _, c := range counts {
		ret /= factorials(c)
	}
	return ret
}
