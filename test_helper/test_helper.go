package test_helper

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

func CollectTestData(params probabilities.DiceRollParameters, measureRoll func([]int) int) *probabilities.Probabilities {
	dices := make([]int, params.DiceCount)
	for i := range dices {
		dices[i] = 1
	}
	total := int(math.Pow(float64(params.DiceSides), float64(params.DiceCount)))
	index := make(map[int]int)
	for i := 0; i < total; i++ {
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
	values := make([]int, maxValue-minValue+1)
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

func extractValues(probs *probabilities.Probabilities) []int {
	values := make([]int, probs.ValuesCount())
	for i := range values {
		values[i] = probs.ValueCount(probs.MinValue() + i)
	}
	return values
}

func CheckProbabilities(t *testing.T, expected, actual *probabilities.Probabilities) {
	assert.Equal(t, expected.MinValue(), actual.MinValue())
	assert.Equal(t, expected.MaxValue(), actual.MaxValue())
	assert.Equal(t, expected.TotalCount(), actual.TotalCount())
	assert.Equal(t, expected.ValuesCount(), actual.ValuesCount())
	assert.Equal(t, extractValues(expected), extractValues(actual))
}
