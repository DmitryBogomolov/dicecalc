package sum_dice

import (
	"math"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

func CalculateProbabilities(params probabilities.DiceRollParameters) (*probabilities.Probabilities, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	min := params.DiceCount
	max := params.DiceCount * params.DiceSides
	factorials := makeFactorials(params.DiceCount)
	totalCount := int(math.Pow(float64(params.DiceSides), float64(params.DiceCount)))
	len := max - min + 1
	values := make([]int, len)
	rolls := makeInitialRolls(params.DiceCount)
	half := len >> 1
	for i := 0; i <= half; i++ {
		values[i] = measureRolls(rolls, factorials)
		rolls = getNextRolls(rolls, params.DiceSides)
	}
	for i := half + 1; i < len; i++ {
		values[i] = values[len-1-i]
	}
	return probabilities.NewProbabilities(min, max, totalCount, values)
}

func makeInitialRolls(diceCount int) []*_Roll {
	dices := make([]byte, diceCount)
	for i := range dices {
		dices[i] = 1
	}
	return []*_Roll{
		{dices: dices},
	}
}

func measureRoll(roll *_Roll, factorials *_Factorials) int {
	n := len(roll.dices)
	counts := make(map[byte]int)
	for _, dice := range roll.dices {
		counts[dice]++
	}
	ret := factorials.get(n)
	for _, c := range counts {
		ret /= factorials.get(c)
	}
	return ret
}

func measureRolls(rolls []*_Roll, factorials *_Factorials) int {
	count := 0
	for _, roll := range rolls {
		k := measureRoll(roll, factorials)
		count += k
	}
	return count
}

func getNextRollsForRoll(roll *_Roll, diceSides int) []*_Roll {
	var rolls []*_Roll
	dices := roll.dices
	for i := 0; i < len(dices); i++ {
		next := dices[i] + 1
		if next <= byte(diceSides) && (i == len(dices)-1 || next <= dices[i+1]) {
			nextDices := append([]byte(nil), dices...)
			nextDices[i] = next
			roll := &_Roll{dices: nextDices}
			rolls = append(rolls, roll)
		}
	}
	return rolls
}

func getNextRolls(rolls []*_Roll, diceSides int) []*_Roll {
	var result []*_Roll
	index := map[string]int{}
	for _, roll := range rolls {
		list := getNextRollsForRoll(roll, diceSides)
		for _, candidate := range list {
			key := candidate.key()
			if _, has := index[key]; !has {
				index[key] = 1
				result = append(result, candidate)
			}
		}
	}
	return result
}

type _Roll struct {
	dices []byte
}

func (roll *_Roll) key() string {
	return string(roll.dices)
}
