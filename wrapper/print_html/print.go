package print_html

import (
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/wrapper/util"
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
	Value       int
	Count       uint64
	Probability string
}

func Print(probs probabilities.Probabilities, title string) []byte {
	var builder strings.Builder
	var data _TemplateData
	data.Title = title
	data.Total = fmt.Sprintf("%d", probs.TotalVariants())
	formatProb := util.GetProbabilityFormatter(probs)
	items := make([]_Item, probs.Count())
	for i := 0; i < probs.Count(); i++ {
		val, count, probability := probs.Item(i)
		items[i] = _Item{
			val,
			count,
			formatProb(probability),
		}
	}
	data.Items = items
	tmpl.Execute(&builder, data)
	return []byte(builder.String())
}
