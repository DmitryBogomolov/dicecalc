package main

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
	Count       int     `json:"count"`
	Probability float64 `json:"probability"`
}

func displayJson(probs probabilities.Probabilities, title string) string {
	var obj _JsonObject
	obj.Title = title
	obj.Total = probs.TotalCount()
	var items []_JsonItem
	for val := probs.MinValue(); val <= probs.MaxValue(); val++ {
		var item _JsonItem
		item.Value = val
		item.Count = probs.ValueCount(val)
		item.Probability = probs.ValueProbability(val)
		items = append(items, item)
	}
	obj.Values = items
	ret, _ := json.MarshalIndent(obj, "", "  ")
	return string(ret)
}