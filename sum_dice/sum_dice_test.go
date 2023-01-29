package sum_dice_test

import (
	"math"
	"strconv"
	"strings"
	"testing"

	. "github.com/DmitryBogomolov/dicecalc/probabilities"
	. "github.com/DmitryBogomolov/dicecalc/sum_dice"
)

func TestX(t *testing.T) {
	checkProbabilities(t, "3d6")
	checkProbabilities(t, "4d6")
	checkProbabilities(t, "5d6")
	checkProbabilities(t, "5d5")
	checkProbabilities(t, "5d6")
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
