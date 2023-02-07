package sum_dice_par

import "github.com/DmitryBogomolov/dicecalc/dice_roller"

func collectAllRolls(value int, roller *dice_roller.DiceRoller) []dice_roller.DiceRoll {
	rootRoll := initDiceRoll(value, roller)
	index := make(map[int]dice_roller.DiceRoll)
	collectAllRollsRecursive(roller, rootRoll, index)
	var rolls []dice_roller.DiceRoll
	for _, roll := range index {
		rolls = append(rolls, roll)
	}
	return rolls
}

func collectAllRollsRecursive(roller *dice_roller.DiceRoller, roll dice_roller.DiceRoll, index map[int]dice_roller.DiceRoll) {
	index[roller.GetRollIdx(roll)] = roll
	for _, r := range getAllSimilarRolls(roller, roll) {
		if _, has := index[roller.GetRollIdx(r)]; !has {
			collectAllRollsRecursive(roller, r, index)
		}
	}
}
