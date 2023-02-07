package sum_dice_base

import "github.com/DmitryBogomolov/dicecalc/dice_roller"

type MeasureRoll func(dice_roller.DiceRoll) int
type MeasureRolls func([]dice_roller.DiceRoll) int

func MakeRollMeasurer(roller *dice_roller.DiceRoller) MeasureRoll {
	factorials := make([]int, roller.DiceCount())
	factorials[0] = 1
	for i := 1; i < len((factorials)); i++ {
		factorials[i] = factorials[(i-1)] * (i + 1)
	}

	return func(roll dice_roller.DiceRoll) int {
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

func MakeRollsMeasurer(roller *dice_roller.DiceRoller) MeasureRolls {
	measure := MakeRollMeasurer(roller)
	return func(rolls []dice_roller.DiceRoll) int {
		c := 0
		for _, roll := range rolls {
			c += measure(roll)
		}
		return c
	}
}
