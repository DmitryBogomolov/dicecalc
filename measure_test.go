package dicecalc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DmitryBogomolov/dicecalc"
	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/sum_dice_seq"
)

func TestVariant1(t *testing.T) {
	start := time.Now()
	dicecalc.CalculateProbabilities(probabilities.DiceRollParameters{DiceSides: 12, DiceCount: 10})
	end := time.Now()
	delta := end.Sub(start).Seconds()
	fmt.Println(delta)
}

func TestVariant2(t *testing.T) {
	start := time.Now()
	sum_dice_seq.CalculateProbabilities(probabilities.DiceRollParameters{DiceSides: 12, DiceCount: 10})
	end := time.Now()
	delta := end.Sub(start).Seconds()
	fmt.Println(delta)
}
