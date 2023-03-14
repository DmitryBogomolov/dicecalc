package print_json

import (
	"encoding/json"
	"fmt"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/wrapper/util"
)

type _JsonObject struct {
	Title  string      `json:"title"`
	Total  uint64      `json:"total"`
	Values []_JsonItem `json:"values"`
}

type _JsonItem struct {
	Value       int    `json:"value"`
	Count       uint64 `json:"count"`
	Probability string `json:"probability"`
}

func Print(probs probabilities.Probabilities, title string) []byte {
	var obj _JsonObject
	obj.Title = title
	obj.Total = probs.TotalVariants()
	items := make([]_JsonItem, probs.Count())
	probFormat := fmt.Sprintf("%%.%df%%%%", util.GetProbabilityPrecision(probs)-2)
	for i := 0; i < probs.Count(); i++ {
		val, count, probability := probs.Item(i)
		items[i] = _JsonItem{
			val,
			count,
			fmt.Sprintf(probFormat, probability*100),
		}
	}
	obj.Values = items
	ret, _ := json.MarshalIndent(obj, "", "  ")
	return ret
}
