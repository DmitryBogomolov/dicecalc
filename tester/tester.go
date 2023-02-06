package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/sum_dice_par"
	"github.com/DmitryBogomolov/dicecalc/sum_dice_seq"
)

func main() {
	params := probabilities.DiceRollParameters{DiceCount: 10, DiceSides: 16}

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go func() {
		testFunc("seq", params, sum_dice_seq.CalculateProbabilities)
		waitGroup.Done()
	}()
	go func() {
		testFunc("par", params, sum_dice_par.CalculateProbabilities)
		waitGroup.Done()
	}()
	waitGroup.Wait()
}

func testFunc(
	tag string,
	params probabilities.DiceRollParameters,
	f func(probabilities.DiceRollParameters) (*probabilities.Probabilities, error),
) {
	start := time.Now()
	f(params)
	end := time.Now()
	diff := end.Sub(start).Seconds()
	fmt.Printf("%s: %.4f\n", tag, diff)
}
