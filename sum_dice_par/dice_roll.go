package sum_dice_par

import "github.com/DmitryBogomolov/dicecalc/dice_roller"

func initDiceRoll(value int, roller *dice_roller.DiceRoller) dice_roller.DiceRoll {
	roll := roller.IdxToRoll(0)
	rest := value - len(roll)
	k := len(roll) - 1
	for rest > 0 {
		val := roller.DiceSides() - 1
		if rest < val {
			val = rest
		}
		rest -= val
		roll[k] += val
		k--
	}
	return roll
}

func getSimilarRoll(roller *dice_roller.DiceRoller, roll dice_roller.DiceRoll, srcIdx, dstIdx int) dice_roller.DiceRoll {
	if dstIdx == srcIdx-1 {
		if (roll[srcIdx] - roll[dstIdx]) < 2 {
			return nil
		}
	} else {
		if (roll[srcIdx] - roll[srcIdx-1]) < 1 {
			return nil
		}
		if (roll[dstIdx+1] - roll[dstIdx]) < 1 {
			return nil
		}
	}
	newRoll := roller.CloneRoll(roll)
	newRoll[srcIdx]--
	newRoll[dstIdx]++
	return newRoll
}

func getAllSimilarRolls(roller *dice_roller.DiceRoller, roll dice_roller.DiceRoll) []dice_roller.DiceRoll {
	var rolls []dice_roller.DiceRoll
	for i := len(roll) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if roll := getSimilarRoll(roller, roll, i, j); roll != nil {
				rolls = append(rolls, roll)
			}
		}
	}
	return rolls
}
