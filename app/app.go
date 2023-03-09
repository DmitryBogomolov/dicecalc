package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DmitryBogomolov/dicecalc/minmax_dice"
	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/sum_dice"
)

type _CalcFunc func(probabilities.DiceRollParameters) (probabilities.Probabilities, error)
type _DisplayFunc func(probabilities.Probabilities, string) string

var modes = map[string]_CalcFunc{
	"sum": sum_dice.CalculateProbabilities,
	"min": minmax_dice.CalculateMinProbabilities,
	"max": minmax_dice.CalculateMaxProbabilities,
}

var outputs = map[string]_DisplayFunc{
	"raw":  displayRaw,
	"json": displayJson,
	"svg":  displaySvg,
}

func main() {
	modeVar := flag.String("mode", "", "operation")
	schemaVar := flag.String("schema", "", "roll schema")
	outputVar := flag.String("output", "", "output format")
	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
		return
	}

	var calcFn _CalcFunc
	var params probabilities.DiceRollParameters
	var err error
	if calcFn, err = parseMode(*modeVar); err != nil {
		fmt.Println(err)
		return
	}
	if params, err = parseRollSchema(*schemaVar); err != nil {
		fmt.Println(err)
		return
	}

	var probs probabilities.Probabilities
	if probs, err = calcFn(params); err != nil {
		fmt.Println(err)
		return
	}

	displayFn := parseOutput(*outputVar)
	title := fmt.Sprintf("Probabilities of %s (%s) rolls", *schemaVar, *modeVar)
	fmt.Println(displayFn(probs, title))
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

func parseMode(mode string) (_CalcFunc, error) {
	if fn, has := modes[mode]; has {
		return fn, nil
	} else {
		return nil, fmt.Errorf("bad mode: %s", mode)
	}
}

func parseOutput(output string) _DisplayFunc {
	if fn, has := outputs[output]; has {
		return fn
	} else {
		return displayRaw
	}
}
