package minmax_dice

import (
	"strconv"
	"strings"
	"testing"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/test_helper"
)

func Test1dX(t *testing.T) {
	checkMaxProbabilities(t, "1d1")
	checkMaxProbabilities(t, "1d5")
	checkMaxProbabilities(t, "1d10")
	checkMaxProbabilities(t, "1d20")

	checkMinProbabilities(t, "1d1")
	checkMinProbabilities(t, "1d5")
	checkMinProbabilities(t, "1d10")
	checkMinProbabilities(t, "1d20")
}

func Test2dX(t *testing.T) {
	checkMaxProbabilities(t, "2d1")
	checkMaxProbabilities(t, "2d4")
	checkMaxProbabilities(t, "2d6")
	checkMaxProbabilities(t, "2d10")
	checkMaxProbabilities(t, "2d20")

	checkMinProbabilities(t, "2d1")
	checkMinProbabilities(t, "2d4")
	checkMinProbabilities(t, "2d6")
	checkMinProbabilities(t, "2d10")
	checkMinProbabilities(t, "2d20")
}

func Test3dX(t *testing.T) {
	checkMaxProbabilities(t, "3d1")
	checkMaxProbabilities(t, "3d4")
	checkMaxProbabilities(t, "3d6")
	checkMaxProbabilities(t, "3d10")
	checkMaxProbabilities(t, "3d20")

	checkMinProbabilities(t, "3d1")
	checkMinProbabilities(t, "3d4")
	checkMinProbabilities(t, "3d6")
	checkMinProbabilities(t, "3d10")
	checkMinProbabilities(t, "3d20")
}

func Test4dX(t *testing.T) {
	checkMaxProbabilities(t, "4d1")
	checkMaxProbabilities(t, "4d4")
	checkMaxProbabilities(t, "4d6")
	checkMaxProbabilities(t, "4d10")
	checkMaxProbabilities(t, "4d16")
	checkMaxProbabilities(t, "4d20")

	checkMinProbabilities(t, "4d1")
	checkMinProbabilities(t, "4d4")
	checkMinProbabilities(t, "4d6")
	checkMinProbabilities(t, "4d10")
	checkMinProbabilities(t, "4d16")
	checkMinProbabilities(t, "4d20")
}

func Test10dX(t *testing.T) {
	checkMaxProbabilities(t, "10d1")
	checkMaxProbabilities(t, "10d6")

	checkMinProbabilities(t, "10d1")
	checkMinProbabilities(t, "10d6")
}

func checkMaxProbabilities(t *testing.T, name string) {
	t.Run(name, func(t *testing.T) {
		parts := strings.Split(name, "d")
		count, _ := strconv.Atoi(parts[0])
		sides, _ := strconv.Atoi(parts[1])
		params := probabilities.DiceRollParameters{DiceCount: count, DiceSides: sides}
		measure := func(roll []int) int {
			max := 1
			for _, dice := range roll {
				if dice > max {
					max = dice
				}
			}
			return max
		}
		expected := test_helper.CollectTestData(params, measure)
		actual, _ := CalculateMaxProbabilities(params)
		test_helper.CheckProbabilities(t, expected, actual)
	})
}

func checkMinProbabilities(t *testing.T, name string) {
	t.Run(name, func(t *testing.T) {
		parts := strings.Split(name, "d")
		count, _ := strconv.Atoi(parts[0])
		sides, _ := strconv.Atoi(parts[1])
		params := probabilities.DiceRollParameters{DiceCount: count, DiceSides: sides}
		measure := func(roll []int) int {
			min := params.DiceSides
			for _, dice := range roll {
				if dice < min {
					min = dice
				}
			}
			return min
		}
		expected := test_helper.CollectTestData(params, measure)
		actual, _ := CalculateMinProbabilities(params)
		test_helper.CheckProbabilities(t, expected, actual)
	})
}
