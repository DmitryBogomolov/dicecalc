package dicecalc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/DmitryBogomolov/dicecalc"
)

func TestCalculateProbabilities(t *testing.T) {
	probabilities, _ := CalculateProbabilities(DiceRollParameters{DiceSides: 6, DiceCount: 4})
	assert.Equal(t, 4, probabilities.MinValue())
	assert.Equal(t, 24, probabilities.MaxValue())
}
