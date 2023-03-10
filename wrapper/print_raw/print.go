package print_raw

import (
	"fmt"
	"math"
	"strings"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"golang.org/x/exp/constraints"
)

func Print(probs probabilities.Probabilities, title string) []byte {
	var builder strings.Builder
	fmt.Fprintln(&builder, title)
	valueSize, countSize, ratioSize := getColumnSizes(probs)
	format := fmt.Sprintf("%%%dd %%%dd %%%d.%df%%%%\n", valueSize, countSize, ratioSize, 4)
	for i := 0; i < probs.Count(); i++ {
		val, count, probability := probs.Item(i)
		fmt.Fprintf(&builder, format, val, count, probability*100)
	}
	return []byte(builder.String())
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
	countSize := getNumberSize(probs.TotalVariants())
	ratioSize := 8
	return valueSize + 1, countSize + 2, ratioSize
}
