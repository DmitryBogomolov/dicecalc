package sum_dice

import (
	"strconv"
	"strings"
	"testing"

	"github.com/DmitryBogomolov/dicecalc/dice_roller"
	"github.com/DmitryBogomolov/dicecalc/test_helper"
)

func Test1dX(t *testing.T) {
	checkProbabilities(t, "1d1")
	checkProbabilities(t, "1d5")
	checkProbabilities(t, "1d10")
	checkProbabilities(t, "1d20")
}

func Test2dX(t *testing.T) {
	checkProbabilities(t, "2d1")
	checkProbabilities(t, "2d4")
	checkProbabilities(t, "2d6")
	checkProbabilities(t, "2d10")
	checkProbabilities(t, "2d20")
}

func Test3dX(t *testing.T) {
	checkProbabilities(t, "3d1")
	checkProbabilities(t, "3d4")
	checkProbabilities(t, "3d6")
	checkProbabilities(t, "3d10")
	checkProbabilities(t, "3d20")
}

func Test4dX(t *testing.T) {
	checkProbabilities(t, "4d1")
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
		params := dice_roller.DiceRollParameters{DiceCount: count, DiceSides: sides}
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
