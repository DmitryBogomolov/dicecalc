package wrapper

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/DmitryBogomolov/dicecalc/minmax_dice"
	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/sum_dice"
	"github.com/DmitryBogomolov/dicecalc/wrapper/print_html"
	"github.com/DmitryBogomolov/dicecalc/wrapper/print_json"
	"github.com/DmitryBogomolov/dicecalc/wrapper/print_raw"
	"github.com/DmitryBogomolov/dicecalc/wrapper/print_svg"
)

type _CalcFunc func(probabilities.DiceRollParameters) (probabilities.Probabilities, error)
type _DisplayFunc func(probabilities.Probabilities, string) []byte

var modes = map[string]_CalcFunc{
	"sum": sum_dice.CalculateProbabilities,
	"min": minmax_dice.CalculateMinProbabilities,
	"max": minmax_dice.CalculateMaxProbabilities,
}

var outputs = map[string]_DisplayFunc{
	"raw":  print_raw.Print,
	"json": print_json.Print,
	"svg":  print_svg.Print,
	"html": print_html.Print,
}

func Process(mode string, schema string, output string) ([]byte, error) {
	var calcFn _CalcFunc
	var params probabilities.DiceRollParameters
	var err error
	if calcFn, err = parseMode(mode); err != nil {
		return nil, err
	}
	if params, err = parseRollSchema(schema); err != nil {
		return nil, err
	}

	var probs probabilities.Probabilities
	if probs, err = calcFn(params); err != nil {
		return nil, err
	}

	var displayFn _DisplayFunc
	if displayFn, err = parseOutput(output); err != nil {
		return nil, err
	}
	title := fmt.Sprintf("Probabilities of %s (%s) rolls", schema, mode)
	ret := displayFn(probs, title)
	return ret, nil
}

func keys[T any](target map[string]T) []string {
	var ret []string
	for key := range target {
		ret = append(ret, key)
	}
	sort.Sort(sort.StringSlice(ret))
	return ret

}

func Modes() []string {
	return keys(modes)
}

func Outputs() []string {
	return keys(outputs)
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

func parseOutput(output string) (_DisplayFunc, error) {
	if fn, has := outputs[output]; has {
		return fn, nil
	} else {
		return nil, fmt.Errorf("bad output: %s", output)
	}
}
