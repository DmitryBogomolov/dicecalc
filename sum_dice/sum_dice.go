package sum_dice

import (
	"math"

	"github.com/DmitryBogomolov/dicecalc/dice_roller"
)

func CalculateProbabilities(params dice_roller.DiceRollParameters) (*dice_roller.Probabilities, error) {
	minVal := params.DiceCount
	maxVal := params.DiceCount * params.DiceSides
	count := maxVal - minVal + 1
	variants := make([]int, count)
	halfCount := int(math.Ceil(0.5 * float64(count)))
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
