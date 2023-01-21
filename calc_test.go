package dicecalc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/DmitryBogomolov/dicecalc"
)

func TestCalculateProbabilities(t *testing.T) {
	probabilities, _ := CalculateProbabilities(DiceRollParameters{DiceCount: 3, DiceSides: 6})
	assert.Equal(t, 3, probabilities.MinValue())
	assert.Equal(t, 18, probabilities.MaxValue())
	assert.Equal(t, 16, probabilities.ValuesCount())
	probs := make([]float64, 16)
	for i := range probs {
		probs[i] = probabilities.ValueProbability(3 + i)
	}
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
	assert.InEpsilonSlice(t, expected, probs, 0.01)
}
