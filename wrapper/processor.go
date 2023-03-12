package wrapper

import (
	"fmt"
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

type ModeFunc func(probabilities.DiceRollParameters) (probabilities.Probabilities, error)
type PrintFunc func(probabilities.Probabilities, string) []byte

var modeNames []string
var modesCache = map[string]ModeFunc{}

var outputNames []string
var outputsCache = map[string]PrintFunc{}

func init() {
	RegisterMode("sum", sum_dice.CalculateProbabilities)
	RegisterMode("min", minmax_dice.CalculateMinProbabilities)
	RegisterMode("max", minmax_dice.CalculateMaxProbabilities)

	RegisterOutput("raw", print_raw.Print)
	RegisterOutput("json", print_json.Print)
	RegisterOutput("html", print_html.Print)
	RegisterOutput("svg", print_svg.Print)
}

func RegisterMode(mode string, fn ModeFunc) {
	if _, has := modesCache[mode]; has {
		panic(fmt.Sprintf("mode %s is already registered", mode))
	}
	if fn == nil {
		panic("fn is nil")
	}
	modeNames = append(modeNames, mode)
	modesCache[mode] = fn
}

func RegisterOutput(output string, fn PrintFunc) {
	if _, has := outputsCache[output]; has {
		panic(fmt.Sprintf("output %s is already registered", output))
	}
	if fn == nil {
		panic("fn is nil")
	}
	outputNames = append(outputNames, output)
	outputsCache[output] = fn
}

func Process(mode string, schema string, output string) ([]byte, error) {
	var calcFn ModeFunc
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

	var displayFn PrintFunc
	if displayFn, err = parseOutput(output); err != nil {
		return nil, err
	}
	title := fmt.Sprintf("Probabilities of %s (%s) rolls", schema, mode)
	ret := displayFn(probs, title)
	return ret, nil
}

func Modes() []string {
	return append([]string(nil), modeNames...)
}

func Outputs() []string {
	return append([]string(nil), outputNames...)
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

func parseMode(mode string) (ModeFunc, error) {
	if fn, has := modesCache[mode]; has {
		return fn, nil
	} else {
		return nil, fmt.Errorf("bad mode: %s", mode)
	}
}

func parseOutput(output string) (PrintFunc, error) {
	if fn, has := outputsCache[output]; has {
		return fn, nil
	} else {
		return nil, fmt.Errorf("bad output: %s", output)
	}
}
