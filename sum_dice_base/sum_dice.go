package sum_dice_base

import "github.com/DmitryBogomolov/dicecalc/dice_roller"

func CalculateProbabilities(
	params dice_roller.DiceRollParameters,
	calculateValues func(int, *dice_roller.DiceRoller) []int,
) (*dice_roller.Probabilities, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	min := params.DiceCount
	max := params.DiceCount * params.DiceSides
	len := max - min + 1
	values := make([]int, len)
	roller := dice_roller.NewRoller(params)
	half := len >> 1
	vals := calculateValues((half + 1), roller)
	copy(values, vals)
	for i := half + 1; i < len; i++ {
		values[i] = values[(len - 1 - i)]
	}
	return dice_roller.NewProbabilities(min, max, roller.TotalRolls(), values)
}
