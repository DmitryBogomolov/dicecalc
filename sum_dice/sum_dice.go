package sum_dice

import (
	"fmt"
	"math"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

// CalculateProbabilities returns probabilities of sums for dice rolls.
//
// For example for 2d6 rolls probability of 2 (1+1) and 12 (6+6) is 1/36,
// probability of 3 (1+2, 2+1) and 11 (5+6, 6+5) is 2/36, etc.
//
// For rolls of n m-sided dices consider n-dimensional cube with side of m.
// That cube consists of smaller m^n cubes with side of 1. Each small cube defines one specific roll.
//
// Plane x_1 + ... + x_n = t (0 < t < n * m) intersects set of small cubes.
// All of them define rolls with the same sum. Numbers of cubes gives probability of that sum.
func CalculateProbabilities(params probabilities.DiceRollParameters) (*probabilities.Probabilities, error) {
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
	minVal := params.DiceCount
	maxVal := params.DiceCount * params.DiceSides
	count := maxVal - minVal + 1
	variants := make([]int, count)
	// For a single dice there is no need for any calculations.
	if params.DiceCount == 1 {
		fillSimpleVariants(variants)
		return probabilities.NewProbabilities(minVal, maxVal, uint64(params.DiceSides), variants)
	}
	// Only half of cube is required to be inspected. The other half is symmetrical.
	halfCount := int(math.Ceil(0.5 * float64(count)))
	// There are n * m segments along big cube diagonal. There are m small cubes in diagonal.
	// So there are n segments alogn each small cube diagonal. And each advance for 1 / (m * n) along
	// big cube corresponds to 1 / n advance along small cube.
	// On each advance along big diagonal take new volume, distribute it along small cubes
	// and find newly filled ones - they belong to plane intersection.
	fillVariants(variants, halfCount, params.DiceCount, params.DiceSides)
	fillSymmetricVariants(variants, halfCount)
	return probabilities.NewProbabilities(minVal, maxVal, uint64(total), variants)
}

func fillSimpleVariants(variants []int) {
	for i := range variants {
		variants[i] = 1
	}
}

func fillVariants(variants []int, count int, diceCount int, diceSides int) {
	// Precalculate small cube volume steps and make slots to keep track of small cubes filling.
	volDiffs, volCounts := prepareDiffsAndCounts(diceCount)
	prevVolume := 0.0
	t := 1
	factor := math.Pow(float64(diceSides), float64(diceCount))
	for i := 0; i < count; i++ {
		// For calculation purpose [0, mn] is mapped to [0, n].
		currVolume := getHyperVolume(float64(t)/float64(diceSides), diceCount)
		diffVolume := (currVolume - prevVolume) * factor
		// Track small cube filling and find newly filled cubes.
		k := distributeVolume(diffVolume, volDiffs, volCounts)
		prevVolume = currVolume
		t++
		variants[i] = k
	}
}

func prepareDiffsAndCounts(diceCount int) ([]float64, []int) {
	volDiffs := make([]float64, diceCount)
	volCounts := make([]int, diceCount-1)
	volDiffs[0] = 1
	for i := 1; i < diceCount; i++ {
		volDiffs[i] = getHyperVolume(float64(diceCount-i), diceCount)
	}
	for i := 0; i < diceCount-1; i++ {
		volDiffs[i] = volDiffs[i] - volDiffs[i+1]
	}
	return volDiffs, volCounts
}

func distributeVolume(volume float64, volDiffs []float64, volCounts []int) int {
	rest := volume
	for i, cnt := range volCounts {
		if cnt > 0 {
			rest -= volDiffs[i] * float64(cnt)
			volCounts[i] = 0
			if i > 0 {
				volCounts[i-1] = cnt
			}
		}
	}
	ratio := rest / volDiffs[len(volDiffs)-1]
	k := int(math.Round(ratio))
	volCounts[len(volCounts)-1] = k
	return k
}

func fillSymmetricVariants(variants []int, half int) {
	count := len(variants)
	for i := count - 1; i >= half; i-- {
		variants[i] = variants[count-i-1]
	}
}
