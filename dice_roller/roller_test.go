package dice_roller_test

import (
	"testing"

	. "github.com/DmitryBogomolov/dicecalc/dice_roller"
	"github.com/stretchr/testify/assert"
)

func TestDiceRoller_GetRollFromIdx(t *testing.T) {
	roller := NewRoller(DiceRollParameters{DiceCount: 3, DiceSides: 6})
	assert.Equal(t, 0, roller.GetRollIdx([]byte{1, 1, 1}))
	assert.Equal(t, 1, roller.GetRollIdx([]byte{1, 1, 2}))
	assert.Equal(t, 5, roller.GetRollIdx([]byte{1, 1, 6}))
	assert.Equal(t, 6, roller.GetRollIdx([]byte{1, 2, 1}))
	assert.Equal(t, 12, roller.GetRollIdx([]byte{1, 3, 1}))
	assert.Equal(t, 179, roller.GetRollIdx([]byte{5, 6, 6}))
	assert.Equal(t, 215, roller.GetRollIdx([]byte{6, 6, 6}))
}

func TestDiceRoller_IdxToRoll(t *testing.T) {
	roller := NewRoller(DiceRollParameters{DiceCount: 3, DiceSides: 6})
	assert.Equal(t, DiceRoll{1, 1, 1}, roller.IdxToRoll(0))
	assert.Equal(t, DiceRoll{1, 1, 2}, roller.IdxToRoll(1))
	assert.Equal(t, DiceRoll{1, 1, 6}, roller.IdxToRoll(5))
	assert.Equal(t, DiceRoll{1, 2, 1}, roller.IdxToRoll(6))
	assert.Equal(t, DiceRoll{1, 3, 1}, roller.IdxToRoll(12))
	assert.Equal(t, DiceRoll{5, 6, 6}, roller.IdxToRoll(179))
	assert.Equal(t, DiceRoll{6, 6, 6}, roller.IdxToRoll(215))
}

func TestDiceRoller_CloneRoll(t *testing.T) {
	roller := NewRoller(DiceRollParameters{DiceCount: 3, DiceSides: 6})
	assert.Equal(t, DiceRoll{1, 2, 3}, roller.CloneRoll(DiceRoll{1, 2, 3}))
}
