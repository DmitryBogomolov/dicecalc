package sum_dice_base

import (
	"github.com/DmitryBogomolov/dicecalc/outdated/dice_roller"
	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

func CalculateProbabilities(
	params probabilities.DiceRollParameters,
	calculateValues func(int, *dice_roller.DiceRoller) []uint64,
) (probabilities.Probabilities, error) {
	roller, err := dice_roller.NewRoller(params)
	if err != nil {
		return nil, err
	}
	min := params.DiceCount
	max := params.DiceCount * params.DiceSides
	len := max - min + 1
	values := make([]uint64, len)
	half := len >> 1
	vals := calculateValues((half + 1), roller)
	copy(values, vals)
	for i := half + 1; i < len; i++ {
		values[i] = values[(len - 1 - i)]
	}
	return probabilities.NewProbabilities(min, max, roller.TotalRolls(), values)
}
