package sum_dice_par

import "github.com/DmitryBogomolov/dicecalc/probabilities"

func collectAllRolls(value int, params probabilities.DiceRollParameters) []*_DiceRoll {
	rootRoll := initDiceRoll(value, params)
	index := make(map[string]*_DiceRoll)
	collectAllRollsRecursive(rootRoll, index)
	var rolls []*_DiceRoll
	for _, roll := range index {
		rolls = append(rolls, roll)
	}
	return rolls
}

func collectAllRollsRecursive(roll *_DiceRoll, index map[string]*_DiceRoll) {
	index[roll.key()] = roll
	for _, r := range roll.getAllSimilarRolls() {
		if _, has := index[r.key()]; !has {
			collectAllRollsRecursive(r, index)
		}
	}
}
