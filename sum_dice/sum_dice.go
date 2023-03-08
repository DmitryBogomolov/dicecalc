package sum_dice

import (
	"math"

	"github.com/DmitryBogomolov/dicecalc/dice_roller"
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
func CalculateProbabilities(params dice_roller.DiceRollParameters) (*dice_roller.Probabilities, error) {
	minVal := params.DiceCount
	maxVal := params.DiceCount * params.DiceSides
	count := maxVal - minVal + 1
	variants := make([]int, count)
	halfCount := int(math.Ceil(0.5 * float64(count)))
	// Only half of cube is required to be inspected. The other half is symmetrical.
	fillVariants(variants, halfCount, params.DiceCount, params.DiceSides)
	fillSymmetricVariants(variants, halfCount)
	total := uint64(math.Pow(float64(params.DiceSides), float64(params.DiceCount)))
	return dice_roller.NewProbabilities(minVal, maxVal, total, variants)
}

func fillVariants(variants []int, count int, diceCount int, diceSides int) {
	volDiffs, volCounts := prepareDiffsAndCounts(diceCount)
	prevVolume := 0.0
	t := 1
	factor := math.Pow(float64(diceSides), float64(diceCount))
	for i := 0; i < count; i++ {
		currVolume := getHyperVolume(float64(t)/float64(diceSides), diceCount)
		diffVolume := (currVolume - prevVolume) * factor
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
