package dicecalc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DmitryBogomolov/dicecalc"
	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/sum_dice"
)

func TestVariant1(t *testing.T) {
	start := time.Now()
	dicecalc.CalculateProbabilities(probabilities.DiceRollParameters{DiceSides: 20, DiceCount: 10})
	end := time.Now()
	delta := end.Sub(start).Seconds()
	fmt.Println(delta)
}

func TestVariant2(t *testing.T) {
	start := time.Now()
	sum_dice.CalculateProbabilities(probabilities.DiceRollParameters{DiceSides: 20, DiceCount: 10})
	sum_dice.CalculateProbabilities(probabilities.DiceRollParameters{DiceSides: 20, DiceCount: 10})
	end := time.Now()
	delta := end.Sub(start).Seconds()
	fmt.Println(delta)
}
