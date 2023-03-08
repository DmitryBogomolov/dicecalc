package max_dice

import "github.com/DmitryBogomolov/dicecalc/probabilities"

// CalculateProbabilities returns probabilities of maxes for dice rolls.
//
// For example for 2d6 rolls probability of 1 (1+1) is 1/36,
// probability of 2 (1+2, 2+1, 2+2) is 3/36, etc.
//
// For rolls of n m-sided dices consider n-dimensional cube with side of m.
// That cube consists of smaller m^n cubes with side of 1. Each small cube defines one specific roll.
//
// ...
func CalculateProbabilities(params probabilities.DiceRollParameters) (probabilities.Probabilities, error) {
	return nil, nil
}
