package max_dice

import (
	"strconv"
	"strings"
	"testing"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
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

func checkProbabilities(t *testing.T, name string) {
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
		actual, _ := CalculateProbabilities(params)
		test_helper.CheckProbabilities(t, expected, actual)
	})
}
