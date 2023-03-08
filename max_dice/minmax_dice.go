package minmax_dice

import (
	"fmt"
	"math"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

// CalculateMaxProbabilities returns probabilities of max for dice rolls.
//
// For example for 2d6 rolls probability of 1 (1+1) is 1/36,
// probability of 2 (1+2, 2+1, 2+2) is 3/36, etc.
//
// For rolls of n m-sided dices consider n-dimensional cube with side of m.
// That cube consists of smaller m^n cubes with side of 1. Each small cube defines one specific roll.
//
// Subcube with side of 1 contains all small cubes for max of 1. Subcube with side of 2 without previous subcube
// contains all small cubes for max of 2. Etc.
func CalculateMaxProbabilities(params probabilities.DiceRollParameters) (probabilities.Probabilities, error) {
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

// CalculateMinProbabilities returns probabilities of min for dice rolls.
//
// For example for 2d6 rolls probability of 6 (6+6) is 1/36,
// probability of 5 (5+6, 6+5, 5+5) is 3/36, etc.
//
// For rolls of n m-sided dices consider n-dimensional cube with side of m.
// That cube consists of smaller m^n cubes with side of 1. Each small cube defines one specific roll.
//
// The idea with subcube is the same with that of "max" case except that subcubes
// grow from the other side of big cube.
func CalculateMinProbabilities(params probabilities.DiceRollParameters) (probabilities.Probabilities, error) {
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
	reverse(variants)
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

func reverse(variants []int) {
	i, j := 0, len(variants)-1
	for i < j {
		variants[i], variants[j] = variants[j], variants[i]
		i++
		j--
	}
}
