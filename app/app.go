package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/DmitryBogomolov/dicecalc/minmax_dice"
	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/sum_dice"
	"golang.org/x/exp/constraints"
)

type _Func func(probabilities.DiceRollParameters) (probabilities.Probabilities, error)

var modes = map[string]_Func{
	"sum": sum_dice.CalculateProbabilities,
	"min": minmax_dice.CalculateMinProbabilities,
	"max": minmax_dice.CalculateMaxProbabilities,
}

func main() {
	modeVar := flag.String("mode", "", "operation")
	schemaVar := flag.String("schema", "", "roll schema")
	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
		return
	}

	var fn _Func
	var params probabilities.DiceRollParameters
	var err error
	if fn, err = parseMode(*modeVar); err != nil {
		fmt.Println(err)
		return
	}
	if params, err = parseRollSchema(*schemaVar); err != nil {
		fmt.Println(err)
		return
	}

	var probs probabilities.Probabilities
	if probs, err = fn(params); err != nil {
		fmt.Println(err)
		return
	}

	title := fmt.Sprintf("Probabilities of %s (%s) rolls", *schemaVar, *modeVar)
	displayProbabilities(probs, title)
}

func parseRollSchema(schema string) (params probabilities.DiceRollParameters, err error) {
	items := strings.Split(strings.ToLower(schema), "d")
	if len(items) != 2 {
		err = fmt.Errorf("bad schema: %s", schema)
		return
	}
	var diceCount int
	if diceCount, err = strconv.Atoi(items[0]); err != nil {
		err = fmt.Errorf("bad schema: %s", schema)
		return
	}
	var diceSides int
	if diceSides, err = strconv.Atoi(items[1]); err != nil {
		err = fmt.Errorf("bad schema: %s", schema)
		return
	}
	params.DiceCount = diceCount
	params.DiceSides = diceSides
	return
}

func parseMode(mode string) (_Func, error) {
	if fn, has := modes[mode]; has {
		return fn, nil
	} else {
		return nil, fmt.Errorf("bad mode: %s", mode)
	}
}

func displayProbabilities(probs probabilities.Probabilities, title string) {
	fmt.Printf("%s / total %d\n", title, probs.TotalCount())
	valueSize, countSize, ratioSize := getColumnSizes(probs)
	format := fmt.Sprintf("%%%dd %%%dd %%%d.%df%%%%\n", valueSize, countSize, ratioSize, 4)
	for val := probs.MinValue(); val <= probs.MaxValue(); val++ {
		count := probs.ValueCount(val)
		ratio := probs.ValueProbability(val) * 100
		fmt.Printf(format, val, count, ratio)
	}
}

func getNumberSize[T constraints.Integer](num T) int {
	return int(math.Ceil(math.Log10(float64(num))))
}

func getColumnSizes(probs probabilities.Probabilities) (int, int, int) {
	minValueSize := getNumberSize(probs.MinValue())
	maxValueSize := getNumberSize(probs.MaxValue())
	valueSize := 0
	if minValueSize > valueSize {
		valueSize = minValueSize
	}
	if maxValueSize > valueSize {
		valueSize = maxValueSize
	}
	countSize := getNumberSize(probs.TotalCount())
	ratioSize := 8
	return valueSize + 1, countSize + 2, ratioSize
}
