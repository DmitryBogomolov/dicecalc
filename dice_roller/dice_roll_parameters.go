package dice_roller

import (
	"fmt"
)

type DiceRollParameters struct {
	DiceSides int
	DiceCount int
}

const MAX_DICE_COUNT = 64
const MAX_DICE_SIDES = 32

func (params DiceRollParameters) Validate() error {
	if params.DiceSides < 1 || params.DiceSides > MAX_DICE_SIDES {
		return fmt.Errorf("bad sides: %d", params.DiceSides)
	}
	if params.DiceCount < 1 || params.DiceCount > MAX_DICE_COUNT {
		return fmt.Errorf("bad dices: %d", params.DiceCount)
	}
	return nil
}
