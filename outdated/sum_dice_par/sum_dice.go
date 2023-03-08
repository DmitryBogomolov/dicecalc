package sum_dice_par

import (
	"sync"

	dr "github.com/DmitryBogomolov/dicecalc/dice_roller"
	"github.com/DmitryBogomolov/dicecalc/outdated/dice_roller"
	"github.com/DmitryBogomolov/dicecalc/outdated/sum_dice_base"
)

func CalculateProbabilities(params dr.DiceRollParameters) (*dr.Probabilities, error) {
	calculateValues := func(k int, roller *dice_roller.DiceRoller) []int {
		calculate := sum_dice_base.MakeDistinctRollsCalculator(roller)
		result := make([]int, k)
		min := roller.DiceCount()
		waigGroup := sync.WaitGroup{}
		waigGroup.Add(k)
		for i := 0; i < k; i++ {
			go func(t int) {
				rolls := collectAllRolls((t + min), roller)
				result[t] = sum_dice_base.CalculateDistinctRolls(rolls, calculate)
				waigGroup.Done()
			}(i)
		}
		waigGroup.Wait()
		return result
	}
	return sum_dice_base.CalculateProbabilities(params, calculateValues)
}

func collectAllRolls(value int, roller *dice_roller.DiceRoller) []dice_roller.DiceRoll {
	rootRoll := initDiceRoll(value, roller)
	index := make(map[uint64]dice_roller.DiceRoll)
	collectAllRollsRecursive(roller, rootRoll, index)
	var rolls []dice_roller.DiceRoll
	for _, roll := range index {
		rolls = append(rolls, roll)
	}
	return rolls
}

func collectAllRollsRecursive(
	roller *dice_roller.DiceRoller, roll dice_roller.DiceRoll, index map[uint64]dice_roller.DiceRoll,
) {
	index[roller.GetRollIdx(roll)] = roll
	for _, r := range getAllSimilarRolls(roller, roll) {
		if _, has := index[roller.GetRollIdx(r)]; !has {
			collectAllRollsRecursive(roller, r, index)
		}
	}
}

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
