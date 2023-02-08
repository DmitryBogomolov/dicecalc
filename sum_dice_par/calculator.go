package sum_dice_par

import (
	"github.com/DmitryBogomolov/dicecalc/dice_roller"
	"github.com/DmitryBogomolov/dicecalc/sum_dice_base"
)

func CalculateProbabilities(params dice_roller.DiceRollParameters) (*dice_roller.Probabilities, error) {
	calculateValues := func(k int, roller *dice_roller.DiceRoller) []int {
		calculate := sum_dice_base.MakeDistinctRollsCalculator(roller)
		result := make([]int, k)
		min := roller.DiceCount()
		for i := 0; i < k; i++ {
			rolls := collectAllRolls((i + min), roller)
			result[i] = sum_dice_base.CalculateDistinctRolls(rolls, calculate)
		}
		return result
	}
	return sum_dice_base.CalculateProbabilities(params, calculateValues)
}
