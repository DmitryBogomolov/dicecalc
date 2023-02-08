package sum_dice_seq

import (
	"github.com/DmitryBogomolov/dicecalc/dice_roller"
	"github.com/DmitryBogomolov/dicecalc/sum_dice_base"
)

func CalculateProbabilities(params dice_roller.DiceRollParameters) (*dice_roller.Probabilities, error) {
	calculateValues := func(k int, roller *dice_roller.DiceRoller) []int {
		calculate := sum_dice_base.MakeDistinctRollsCalculator(roller)
		rolls := []dice_roller.DiceRoll{roller.IdxToRoll(0)}
		result := make([]int, k)
		for i := 0; i < k; i++ {
			result[i] = sum_dice_base.CalculateDistinctRolls(rolls, calculate)
			rolls = getNextRolls(rolls, roller)
		}
		return result
	}
	return sum_dice_base.CalculateProbabilities(params, calculateValues)
}

func getNextRollsForRoll(roll dice_roller.DiceRoll, roller *dice_roller.DiceRoller) []dice_roller.DiceRoll {
	var rolls []dice_roller.DiceRoll
	for i := 0; i < len(roll); i++ {
		next := roll[i] + 1
		threshold := roller.DiceSides()
		if i < (len(roll) - 1) {
			threshold = roll[(i + 1)]
		}
		if next <= threshold {
			nextRoll := roller.CloneRoll(roll)
			nextRoll[i] = next
			rolls = append(rolls, nextRoll)
		}
	}
	return rolls
}

func getNextRolls(rolls []dice_roller.DiceRoll, roller *dice_roller.DiceRoller) []dice_roller.DiceRoll {
	var result []dice_roller.DiceRoll
	index := map[uint64]int{}
	for _, roll := range rolls {
		list := getNextRollsForRoll(roll, roller)
		for _, candidate := range list {
			key := roller.GetRollIdx(candidate)
			if _, has := index[key]; !has {
				index[key] = 1
				result = append(result, candidate)
			}
		}
	}
	return result
}
