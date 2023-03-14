package print_csv

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

func Print(probs probabilities.Probabilities, title string) []byte {
	var builder strings.Builder
	writer := csv.NewWriter(&builder)
	writer.Write([]string{"value", "count", "probability"})
	for i := 0; i < probs.Count(); i++ {
		val, count, probability := probs.Item(i)
		writer.Write([]string{fmt.Sprintf("%d", val), fmt.Sprintf("%d", count), fmt.Sprintf("%.2f%%", probability)})
	}
	writer.Flush()
	// var obj _JsonObject
	// obj.Title = title
	// obj.Total = probs.TotalVariants()
	// var items []_JsonItem
	// for i := 0; i < probs.Count(); i++ {
	// 	val, count, probability := probs.Item(i)
	// 	var item _JsonItem
	// 	item.Value = val
	// 	item.Count = count
	// 	item.Probability = probability
	// 	items = append(items, item)
	// }
	// obj.Values = items
	// ret, _ := json.MarshalIndent(obj, "", "  ")
	return []byte(builder.String())
}
