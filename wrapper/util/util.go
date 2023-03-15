package util

import (
	"fmt"
	"math"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"golang.org/x/exp/constraints"
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

func GetNumberSize[T constraints.Integer](num T) int {
	return int(math.Ceil(math.Log10(float64(num))))
}

func GetColumnSizes(probs probabilities.Probabilities, formatProb func(float64) string) (int, int, int) {
	minValueSize := GetNumberSize(probs.MinValue())
	maxValueSize := GetNumberSize(probs.MaxValue())
	valueSize := 0
	if minValueSize > valueSize {
		valueSize = minValueSize
	}
	if maxValueSize > valueSize {
		valueSize = maxValueSize
	}
	countSize := GetNumberSize(probs.TotalVariants())
	ratioSize := len(formatProb(1.0))
	return valueSize + 1, countSize + 2, ratioSize + 2
}
