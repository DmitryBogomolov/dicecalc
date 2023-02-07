package sum_dice_par

import (
	"github.com/DmitryBogomolov/dicecalc/dice_roller"
	"github.com/DmitryBogomolov/dicecalc/sum_dice_base"
)

func CalculateProbabilities(params dice_roller.DiceRollParameters) (*dice_roller.Probabilities, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	min := params.DiceCount
	max := params.DiceCount * params.DiceSides
	roller := dice_roller.NewRoller(params)
	measureRolls := sum_dice_base.MakeRollsMeasurer(roller)
	len := max - min + 1
	values := make([]int, len)
	half := len >> 1
	for i := 0; i <= half; i++ {
		k := i
		rolls := collectAllRolls(k+min, roller)
		values[k] = measureRolls(rolls)
	}
	for i := half + 1; i < len; i++ {
		values[i] = values[(len - 1 - i)]
	}
	return dice_roller.NewProbabilities(min, max, roller.TotalRolls(), values)
}
