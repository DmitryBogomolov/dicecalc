package sum_dice_par_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/DmitryBogomolov/dicecalc/dice_roller"
	. "github.com/DmitryBogomolov/dicecalc/sum_dice_par"
	"github.com/DmitryBogomolov/dicecalc/test_helper"
)

func TestCalculateProbabilities(t *testing.T) {
	probabilities, _ := CalculateProbabilities(DiceRollParameters{DiceCount: 3, DiceSides: 6})
	assert.Equal(t, 3, probabilities.MinValue())
	assert.Equal(t, 18, probabilities.MaxValue())
	assert.Equal(t, 16, probabilities.ValuesCount())
}

func TestValidation(t *testing.T) {
	{
		probabilities, err := CalculateProbabilities(DiceRollParameters{DiceCount: 0, DiceSides: 4})
		assert.Nil(t, probabilities)
		assert.Error(t, err)
	}
	{
		probabilities, err := CalculateProbabilities(DiceRollParameters{DiceCount: 3, DiceSides: 0})
		assert.Nil(t, probabilities)
		assert.Error(t, err)
	}
	{
		probabilities, err := CalculateProbabilities(DiceRollParameters{DiceCount: MAX_DICE_COUNT + 1, DiceSides: 4})
		assert.Nil(t, probabilities)
		assert.Error(t, err)
	}
	{
		probabilities, err := CalculateProbabilities(DiceRollParameters{DiceCount: 3, DiceSides: MAX_DICE_SIDES + 1})
		assert.Nil(t, probabilities)
		assert.Error(t, err)
	}
}

func Test3dX(t *testing.T) {
	checkProbabilities(t, "3d4")
	checkProbabilities(t, "3d6")
	checkProbabilities(t, "3d10")
	checkProbabilities(t, "3d20")
}

func Test4dX(t *testing.T) {
	checkProbabilities(t, "4d4")
	checkProbabilities(t, "4d6")
	checkProbabilities(t, "4d10")
	checkProbabilities(t, "4d16")
	checkProbabilities(t, "4d20")
}

func Test10dX(t *testing.T) {
	checkProbabilities(t, "10d1")
	checkProbabilities(t, "10d6")
}

func checkProbabilities(t *testing.T, name string) {
	t.Run(name, func(t *testing.T) {
		parts := strings.Split(name, "d")
		count, _ := strconv.Atoi(parts[0])
		sides, _ := strconv.Atoi(parts[1])
		params := DiceRollParameters{DiceCount: count, DiceSides: sides}
		measure := func(roll []int) int {
			sum := 0
			for _, dice := range roll {
				sum += dice
			}
			return sum
		}
		expected := test_helper.CollectTestData(params, measure)
		actual, _ := CalculateProbabilities(params)
		test_helper.CheckProbabilities(t, expected, actual)
	})
}
