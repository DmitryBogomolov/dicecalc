package dice_roller

import "math"

type DiceRoller struct {
	diceCount  int
	diceSides  int
	totalRolls uint64
}

type DiceRoll []int

func NewRoller(params DiceRollParameters) *DiceRoller {
	return &DiceRoller{
		diceCount:  params.DiceCount,
		diceSides:  params.DiceSides,
		totalRolls: uint64(math.Pow(float64(params.DiceSides), float64(params.DiceCount))),
	}
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
