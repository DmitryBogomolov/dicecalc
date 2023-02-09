package dice_roller

import (
	"fmt"
	"math"
)

type DiceRoller struct {
	diceCount  int
	diceSides  int
	totalRolls uint64
}

type DiceRoll []int

func NewRoller(params DiceRollParameters) (roller *DiceRoller, err error) {
	if params.DiceCount < 1 {
		return nil, fmt.Errorf("bad dice count: %d", params.DiceCount)
	}
	if params.DiceSides < 1 {
		return nil, fmt.Errorf("bad dice sides: %d", params.DiceSides)
	}
	total := math.Pow(float64(params.DiceSides), float64(params.DiceCount))
	if total > math.MaxUint64 {
		return nil, fmt.Errorf("too big values: %d or %d", params.DiceCount, params.DiceSides)
	}
	return &DiceRoller{
		diceCount:  params.DiceCount,
		diceSides:  params.DiceSides,
		totalRolls: uint64(total),
	}, nil
}

func (roller *DiceRoller) DiceCount() int {
	return roller.diceCount
}

func (roller *DiceRoller) DiceSides() int {
	return roller.diceSides
}

func (roller *DiceRoller) TotalRolls() uint64 {
	return roller.totalRolls
}

func (roller *DiceRoller) GetRollIdx(roll DiceRoll) uint64 {
	idx := uint64(0)
	for _, dice := range roll {
		idx = (uint64(roller.diceSides) * idx) + uint64(dice) - 1
	}
	return idx
}

func (roller *DiceRoller) IdxToRoll(idx uint64) DiceRoll {
	dices := make(DiceRoll, roller.diceCount)
	divisor := roller.totalRolls
	residue := idx
	for i := 0; i < roller.diceCount; i++ {
		divisor /= uint64(roller.diceSides)
		dices[i] = int((residue / divisor) + 1)
		residue %= divisor
	}
	return dices
}

func (roller *DiceRoller) CloneRoll(roll DiceRoll) DiceRoll {
	return append(DiceRoll(nil), roll...)
}
