package sum_dice_base

import (
	"github.com/DmitryBogomolov/dicecalc/outdated/dice_roller"
)

type DistinctRollsCalculator func(dice_roller.DiceRoll) uint64

func MakeDistinctRollsCalculator(roller *dice_roller.DiceRoller) DistinctRollsCalculator {
	factorials := make([]uint64, roller.DiceCount())
	factorials[0] = 1
	for i := 1; i < len(factorials); i++ {
		factorials[i] = factorials[(i-1)] * uint64(i+1)
	}

	return func(roll dice_roller.DiceRoll) uint64 {
		n := len(roll)
		counts := make(map[int]int)
		for _, dice := range roll {
			counts[dice]++
		}
		ret := factorials[(n - 1)]
		for _, c := range counts {
			ret /= factorials[(c - 1)]
		}
		return ret
	}
}

func CalculateDistinctRolls(rolls []dice_roller.DiceRoll, calculate DistinctRollsCalculator) uint64 {
	count := uint64(0)
	for _, roll := range rolls {
		count += calculate(roll)
	}
	return count
}
