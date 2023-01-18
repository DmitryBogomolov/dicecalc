package dicecalc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/DmitryBogomolov/dicecalc"
)

func TestCalculateProbabilities(t *testing.T) {
	probabilities := CalculateProbabilities(DiceSchema{Sides: 6, Count: 4})
	assert.Equal(t, 4, probabilities.Min())
	assert.Equal(t, 24, probabilities.Max())
}
