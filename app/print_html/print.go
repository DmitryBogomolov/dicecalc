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
	Items []_Item
}
type _Item struct {
	Value       string
	Count       string
	Probability string
}

func Print(probs probabilities.Probabilities, title string) string {
	var builder strings.Builder
	var data _TemplateData
	data.Title = title
	var items []_Item
	for val := probs.MinValue(); val <= probs.MaxValue(); val++ {
		var item _Item
		item.Value = fmt.Sprintf("%d", val)
		item.Count = fmt.Sprintf("%d", probs.ValueVariants(val))
		item.Probability = fmt.Sprintf("%.2f%%", probs.ValueProbability(val)*100)
		items = append(items, item)
	}
	data.Items = items
	tmpl.Execute(&builder, data)
	return builder.String()
}
