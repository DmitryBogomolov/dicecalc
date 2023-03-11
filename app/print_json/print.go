package print_json

import (
	"encoding/json"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

type _JsonObject struct {
	Title  string      `json:"title"`
	Total  uint64      `json:"total"`
	Values []_JsonItem `json:"values"`
}

type _JsonItem struct {
	Value       int     `json:"value"`
	Count       uint64  `json:"count"`
	Probability float64 `json:"probability"`
}

func Print(probs probabilities.Probabilities, title string) string {
	var obj _JsonObject
	obj.Title = title
	obj.Total = probs.TotalVariants()
	var items []_JsonItem
	for i := 0; i < probs.Count(); i++ {
		val, count, probability := probs.Item(i)
		var item _JsonItem
		item.Value = val
		item.Count = count
		item.Probability = probability
		items = append(items, item)
	}
	obj.Values = items
	ret, _ := json.MarshalIndent(obj, "", "  ")
	return string(ret)
}
