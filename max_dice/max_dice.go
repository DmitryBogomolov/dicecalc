package max_dice

import (
	"fmt"
	"math"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

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
	if params.DiceCount < 1 {
		return nil, fmt.Errorf("bad dice count: %d", params.DiceCount)
	}
	if params.DiceSides < 1 {
		return nil, fmt.Errorf("bad dice sides: %d", params.DiceSides)
	}
	total := math.Pow(float64(params.DiceSides), float64(params.DiceCount))
	if total > math.MaxUint64 {
		return nil, fmt.Errorf("too big values: %d^%d", params.DiceCount, params.DiceSides)
	}
	minVal := 1
	maxVal := params.DiceSides
	count := maxVal - minVal + 1
	variants := make([]int, count)
	fillVariants(variants, params.DiceCount, params.DiceSides)
	return probabilities.NewProbabilities(minVal, maxVal, uint64(total), variants)
}

func fillVariants(variants []int, diceCount int, diceSides int) {
	prevCount := uint64(0)
	for i := 0; i < len(variants); i++ {
		currCount := uint64(math.Pow(float64(i+1), float64(diceCount)))
		variants[i] = int(currCount - prevCount)
		prevCount = currCount
	}
}
