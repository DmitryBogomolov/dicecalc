package test_helper

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

func CollectTestData(params probabilities.DiceRollParameters, measureRoll func([]int) int) probabilities.Probabilities {
	dices := make([]int, params.DiceCount)
	for i := range dices {
		dices[i] = 1
	}
	total := uint64(math.Pow(float64(params.DiceSides), float64(params.DiceCount)))
	index := make(map[int]uint64)
	for i := uint64(0); i < total; i++ {
		value := measureRoll(dices)
		advanceRoll(dices, params.DiceSides)
		index[value]++
	}
	minValue := math.MaxInt32
	maxValue := 0
	for key := range index {
		if key < minValue {
			minValue = key
		}
		if key > maxValue {
			maxValue = key
		}
	}
	values := make([]uint64, maxValue-minValue+1)
	for key, val := range index {
		values[key-minValue] = val
	}
	probs, _ := probabilities.NewProbabilities(minValue, maxValue, total, values)
	return probs
}

func advanceRoll(dices []int, diceSides int) {
	for idx := len(dices) - 1; idx >= 0; idx-- {
		dices[idx]++
		if dices[idx] > diceSides {
			dices[idx] = 1
		} else {
			break
		}
	}
}

func extractValues(probs probabilities.Probabilities) []uint64 {
	var values []uint64
	for i := 0; i < probs.Count(); i++ {
		_, val, _ := probs.Item(i)
		values = append(values, val)
	}
	return values
}

func CheckProbabilities(t *testing.T, expected, actual probabilities.Probabilities) {
	assert.Equal(t, expected.MinValue(), actual.MinValue())
	assert.Equal(t, expected.MaxValue(), actual.MaxValue())
	assert.Equal(t, expected.TotalVariants(), actual.TotalVariants())
	assert.Equal(t, expected.Count(), actual.Count())
	assert.Equal(t, extractValues(expected), extractValues(actual))
}
