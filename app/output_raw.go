package main

import (
	"fmt"
	"math"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"golang.org/x/exp/constraints"
)

func displayRaw(probs probabilities.Probabilities, title string) {
	fmt.Printf("%s / total %d\n", title, probs.TotalCount())
	valueSize, countSize, ratioSize := getColumnSizes(probs)
	format := fmt.Sprintf("%%%dd %%%dd %%%d.%df%%%%\n", valueSize, countSize, ratioSize, 4)
	for val := probs.MinValue(); val <= probs.MaxValue(); val++ {
		count := probs.ValueCount(val)
		ratio := probs.ValueProbability(val) * 100
		fmt.Printf(format, val, count, ratio)
	}
}

func getNumberSize[T constraints.Integer](num T) int {
	return int(math.Ceil(math.Log10(float64(num))))
}

func getColumnSizes(probs probabilities.Probabilities) (int, int, int) {
	minValueSize := getNumberSize(probs.MinValue())
	maxValueSize := getNumberSize(probs.MaxValue())
	valueSize := 0
	if minValueSize > valueSize {
		valueSize = minValueSize
	}
	if maxValueSize > valueSize {
		valueSize = maxValueSize
	}
	countSize := getNumberSize(probs.TotalCount())
	ratioSize := 8
	return valueSize + 1, countSize + 2, ratioSize
}
