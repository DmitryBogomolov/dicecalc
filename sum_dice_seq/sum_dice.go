package sum_dice_seq

import (
	"github.com/DmitryBogomolov/dicecalc/dice_roller"
)

func CalculateProbabilities(params dice_roller.DiceRollParameters) (*dice_roller.Probabilities, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	min := params.DiceCount
	max := params.DiceCount * params.DiceSides
	factorials := dice_roller.NewFactorials(params.DiceCount)
	len := max - min + 1
	values := make([]int, len)
	roller := dice_roller.NewRoller(params)
	rolls := []dice_roller.DiceRoll{roller.IdxToRoll(0)}
	half := len >> 1
	for i := 0; i <= half; i++ {
		values[i] = measureRolls(rolls, factorials)
		rolls = getNextRolls(rolls, roller)
	}
	for i := half + 1; i < len; i++ {
		values[i] = values[(len - 1 - i)]
	}
	return dice_roller.NewProbabilities(min, max, roller.TotalRolls(), values)
}

func measureRoll(roll dice_roller.DiceRoll, factorials *dice_roller.Factorials) int {
	n := len(roll)
	counts := make(map[int]int)
	for _, dice := range roll {
		counts[dice]++
	}
	ret := factorials.Get(n)
	for _, c := range counts {
		ret /= factorials.Get(c)
	}
	return ret
}

func measureRolls(rolls []dice_roller.DiceRoll, factorials *dice_roller.Factorials) int {
	count := 0
	for _, roll := range rolls {
		k := measureRoll(roll, factorials)
		count += k
	}
	return count
}

func getNextRollsForRoll(roll dice_roller.DiceRoll, roller *dice_roller.DiceRoller) []dice_roller.DiceRoll {
	var rolls []dice_roller.DiceRoll
	for i := 0; i < len(roll); i++ {
		next := roll[i] + 1
		threshold := roller.DiceSides()
		if i < (len(roll) - 1) {
			threshold = roll[(i + 1)]
		}
		if next <= threshold {
			nextRoll := roller.CloneRoll(roll)
			nextRoll[i] = next
			rolls = append(rolls, nextRoll)
		}
	}
	return rolls
}

func getNextRolls(rolls []dice_roller.DiceRoll, roller *dice_roller.DiceRoller) []dice_roller.DiceRoll {
	var result []dice_roller.DiceRoll
	index := map[int]int{}
	for _, roll := range rolls {
		list := getNextRollsForRoll(roll, roller)
		for _, candidate := range list {
			key := roller.GetRollIdx(candidate)
			if _, has := index[key]; !has {
				index[key] = 1
				result = append(result, candidate)
			}
		}
	}
	return result
}
