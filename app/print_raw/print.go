package print_raw

import (
	"fmt"
	"math"
	"strings"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"golang.org/x/exp/constraints"
)

func Print(probs probabilities.Probabilities, title string) string {
	var builder strings.Builder
	builder.WriteString(title + "\n")
	valueSize, countSize, ratioSize := getColumnSizes(probs)
	format := fmt.Sprintf("%%%dd %%%dd %%%d.%df%%%%\n", valueSize, countSize, ratioSize, 4)
	for val := probs.MinValue(); val <= probs.MaxValue(); val++ {
		count := probs.ValueVariants(val)
		ratio := probs.ValueProbability(val) * 100
		builder.WriteString(fmt.Sprintf(format, val, count, ratio))
	}
	return builder.String()
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
	countSize := getNumberSize(probs.VariantsCount())
	ratioSize := 8
	return valueSize + 1, countSize + 2, ratioSize
}
