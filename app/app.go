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
	displayRaw(probs, title)
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
