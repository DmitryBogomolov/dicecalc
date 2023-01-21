package dicecalc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/DmitryBogomolov/dicecalc"
)

func makeProbabilitiesList(probabilities *Probabilities) []float64 {
	probs := make([]float64, probabilities.ValuesCount())
	for i := range probs {
		probs[i] = probabilities.ValueProbability(i + probabilities.MinValue())
	}
	return probs
}

func TestCalculateProbabilities(t *testing.T) {
	probabilities, _ := CalculateProbabilities(DiceRollParameters{DiceCount: 3, DiceSides: 6})
	assert.Equal(t, 3, probabilities.MinValue())
	assert.Equal(t, 18, probabilities.MaxValue())
	assert.Equal(t, 16, probabilities.ValuesCount())
	actual := makeProbabilitiesList(probabilities)
	expected := []float64{
		0.0046,
		0.0139,
		0.0278,
		0.0463,
		0.0694,
		0.0972,
		0.1157,
		0.1250,
		0.1250,
		0.1157,
		0.0972,
		0.0694,
		0.0463,
		0.0278,
		0.0139,
		0.0046,
	}
	assert.InEpsilonSlice(t, expected, actual, 0.01)
}
