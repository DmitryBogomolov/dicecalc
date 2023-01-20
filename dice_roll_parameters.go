package dicecalc

import (
	"fmt"
	"math"
)

type DiceRollParameters struct {
	DiceSides int
	DiceCount int
}

const MAX_DICE_COUNT = 64
const MAX_DICE_SIDES = 32

func validateParameters(params DiceRollParameters) error {
	if params.DiceSides < 1 || params.DiceSides > MAX_DICE_SIDES {
		return fmt.Errorf("bad sides: %d", params.DiceSides)
	}
	if params.DiceCount < 1 || params.DiceCount > MAX_DICE_COUNT {
		return fmt.Errorf("bad dices: %d", params.DiceCount)
	}
	return nil
}

func getValueRange(params DiceRollParameters) (int, int) {
	return params.DiceCount, params.DiceCount * params.DiceSides
}

func getVariantsCount(params DiceRollParameters) int {
	return int(math.Pow(float64(params.DiceSides), float64(params.DiceCount)))
}
