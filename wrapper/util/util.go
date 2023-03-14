package util

import (
	"math"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

func GetProbabilityPrecision(probs probabilities.Probabilities) int {
	p := math.Ceil(math.Log10(float64(probs.TotalVariants())))
	return int(p)
}
