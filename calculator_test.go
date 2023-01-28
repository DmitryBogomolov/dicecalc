package dicecalc_test

import (
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/DmitryBogomolov/dicecalc"
	. "github.com/DmitryBogomolov/dicecalc/probabilities"
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

func measureRoll(dices []int) int {
	sum := 0
	for _, dice := range dices {
		sum += dice
	}
	return sum
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

func makeExpectedData(diceCount, diceSides int) []float64 {
	dices := make([]int, diceCount)
	total := 1
	for i := range dices {
		dices[i] = 1
		total *= diceSides
	}
	counter := make([]int, diceCount*diceSides-diceCount+1)
	for i := 0; i < total; i++ {
		value := measureRoll(dices)
		advanceRoll(dices, diceSides)
		counter[value-diceCount]++
	}
	data := make([]float64, len(counter))
	for i := range counter {
		data[i] = float64(counter[i]) / float64(total)
	}
	return data
}

func makeActualData(diceCount, diceSides int) []float64 {
	probabilies, _ := CalculateProbabilities(DiceRollParameters{DiceCount: diceCount, DiceSides: diceSides})
	data := make([]float64, probabilies.ValuesCount())
	for i := range data {
		data[i] = probabilies.ValueProbability(probabilies.MinValue() + i)
	}
	return data
}

func checkProbabilities(t *testing.T, name string) {
	t.Run(name, func(t *testing.T) {
		parts := strings.Split(name, "d")
		count, _ := strconv.Atoi(parts[0])
		sides, _ := strconv.Atoi(parts[1])
		expected := makeExpectedData(count, sides)
		actual := makeActualData(count, sides)
		for i := range expected {
			if math.Abs(actual[i]-expected[i]) > 0.0001 {
				t.Errorf("%dth element: expected %f actual %f", i, expected[i], actual[i])
			}
		}
	})
}
