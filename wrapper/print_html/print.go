package print_html

import (
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

//go:embed template.html
var tmplStr string

var tmpl = template.Must(template.New("svg").Parse(tmplStr))

type _TemplateData struct {
	Title string
	Total string
	Items []_Item
}
type _Item struct {
	Value       string
	Count       string
	Probability string
}

func Print(probs probabilities.Probabilities, title string) []byte {
	var builder strings.Builder
	var data _TemplateData
	data.Title = title
	data.Total = fmt.Sprintf("%d", probs.TotalVariants())
	var items []_Item
	for i := 0; i < probs.Count(); i++ {
		val, count, probability := probs.Item(i)
		var item _Item
		item.Value = fmt.Sprintf("%d", val)
		item.Count = fmt.Sprintf("%d", count)
		item.Probability = fmt.Sprintf("%.2f%%", probability*100)
		items = append(items, item)
	}
	data.Items = items
	tmpl.Execute(&builder, data)
	return []byte(builder.String())
}
