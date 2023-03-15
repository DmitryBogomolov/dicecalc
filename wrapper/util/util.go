package util

import (
	"fmt"
	"math"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

func getProbabilityPrecision(probs probabilities.Probabilities) int {
	p := math.Ceil(math.Log10(float64(probs.TotalVariants())))
	return int(p)
}

func GetProbabilityFormatter(probs probabilities.Probabilities) func(float64) string {
	prec := getProbabilityPrecision(probs)
	k := prec - 2
	if k < 1 {
		k = 1
	}
	format := fmt.Sprintf("%%.%df%%%%", k)
	return func(val float64) string {
		return fmt.Sprintf(format, val*100)
	}
}
