package dice_roller

import "math"

type DiceRoller struct {
	diceCount  int
	diceSides  int
	totalRolls int
}

type DiceRoll []byte

func NewRoller(params DiceRollParameters) *DiceRoller {
	return &DiceRoller{
		diceCount:  params.DiceCount,
		diceSides:  params.DiceSides,
		totalRolls: int(math.Pow(float64(params.DiceSides), float64(params.DiceCount))),
	}
}

func (roller *DiceRoller) DiceCount() int {
	return roller.diceCount
}

func (roller *DiceRoller) DiceSides() int {
	return roller.diceSides
}

func (roller *DiceRoller) TotalRolls() int {
	return roller.totalRolls
}

func (roller *DiceRoller) GetRollIdx(roll DiceRoll) int {
	idx := 0
	for _, dice := range roll {
		idx = (roller.diceSides * idx) + int(dice) - 1
	}
	return idx
}

func (roller *DiceRoller) IdxToRoll(idx int) DiceRoll {
	dices := make(DiceRoll, roller.diceCount)
	divisor := roller.totalRolls
	residue := idx
	for i := 0; i < roller.diceCount; i++ {
		divisor /= roller.diceSides
		dices[i] = byte(residue/divisor) + 1
		residue %= divisor
	}
	return dices
}

func (roller *DiceRoller) CloneRoll(roll DiceRoll) DiceRoll {
	return append(DiceRoll(nil), roll...)
}

func (roller *DiceRoller) IsValidDice(dice byte) bool {
	return 1 <= dice && dice <= byte(roller.diceSides)
}