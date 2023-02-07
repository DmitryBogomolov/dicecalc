package sum_dice_par

import "github.com/DmitryBogomolov/dicecalc/dice_roller"

func collectAllRolls(value int, roller *dice_roller.DiceRoller) []dice_roller.DiceRoll {
	rootRoll := initDiceRoll(value, roller)
	index := make(map[string]dice_roller.DiceRoll)
	collectAllRollsRecursive(roller, rootRoll, index)
	var rolls []dice_roller.DiceRoll
	for _, roll := range index {
		rolls = append(rolls, roll)
	}
	return rolls
}

func collectAllRollsRecursive(roller *dice_roller.DiceRoller, roll dice_roller.DiceRoll, index map[string]dice_roller.DiceRoll) {
	index[rollKey(roll)] = roll
	for _, r := range getAllSimilarRolls(roller, roll) {
		if _, has := index[rollKey(r)]; !has {
			collectAllRollsRecursive(roller, r, index)
		}
	}
}
