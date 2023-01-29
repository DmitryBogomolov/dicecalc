package dicecalc

import (
	"math"
	"sync"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

func CalculateProbabilities(params probabilities.DiceRollParameters) (*probabilities.Probabilities, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	min := params.DiceCount
	max := params.DiceCount * params.DiceSides
	totalCount := int(math.Pow(float64(params.DiceSides), float64(params.DiceCount)))
	factorials := makeFactorials(params.DiceCount)
	len := max - min + 1
	values := make([]int, len)
	half := len >> 1
	var wg sync.WaitGroup
	for i := 0; i <= half; i++ {
		wg.Add(1)
		go func(k int) {
			rolls := collectAllRolls(k+min, params)
			values[k] = calculateValueSlots(rolls, factorials)
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := half + 1; i < len; i++ {
		values[i] = values[len-1-i]
	}
	return probabilities.NewProbabilities(min, max, totalCount, values)
}

func calculateValueSlots(rolls []*_DiceRoll, factorials *_Factorials) int {
	count := 0
	for _, roll := range rolls {
		k := calculateRollCount(roll, factorials)
		count += k
	}
	return count
}

func calculateRollCount(roll *_DiceRoll, factorials *_Factorials) int {
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
