package sum_dice

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DmitryBogomolov/dicecalc/dice_roller"
)

func TestCalculateProbabilities(t *testing.T) {
	probs, _ := CalculateProbabilities(dice_roller.DiceRollParameters{DiceCount: 2, DiceSides: 4})
	assert.Equal(t, 2, probs.MinValue())
	assert.Equal(t, 12, probs.MaxValue())
}
