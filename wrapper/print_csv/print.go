package print_csv

import (
	"fmt"
	"strings"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/wrapper/util"
)

func Print(probs probabilities.Probabilities, title string) []byte {
	var builder strings.Builder
	fmt.Fprintf(&builder, "# %s\n", title)
	formatProb := util.GetProbabilityFormatter(probs)
	valSize, countSize, probSize := util.GetColumnSizes(probs, formatProb)
	format := fmt.Sprintf("%%%dd , %%%dd , %%%ds\n", valSize, countSize, probSize)
	for i := 0; i < probs.Count(); i++ {
		val, count, probability := probs.Item(i)
		fmt.Fprintf(&builder, format, val, count, formatProb(probability))
	}
	fmt.Fprintf(&builder, "# Total count: %d\n", probs.TotalVariants())
	return []byte(builder.String())
}
