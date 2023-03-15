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
	format := fmt.Sprintf("%%d,%%d,%%.%df%%%%\n", util.GetProbabilityPrecision(probs))
	for i := 0; i < probs.Count(); i++ {
		val, count, probability := probs.Item(i)
		fmt.Fprintf(&builder, format, val, count, probability*100)
	}
	fmt.Fprintf(&builder, "# Total count: %d\n", probs.TotalVariants())
	return []byte(builder.String())
}
