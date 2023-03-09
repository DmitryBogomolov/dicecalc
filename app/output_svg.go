package main

import (
	"html/template"
	"strings"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

var tmplStr = `<svg width="{{ .Width }}" height="{{ .Height }}" xmlns="http://www.w3.org/2000/svg">
  <text x="20" y="40">Hello World</text>
</svg>`

var tmpl = template.Must(template.New("svg").Parse(tmplStr))

type _TmplContent struct {
	Width  int
	Height int
}

func displaySvg(probs probabilities.Probabilities, title string) string {
	var builder strings.Builder
	var ctx _TmplContent
	ctx.Width = 640
	ctx.Height = 480
	tmpl.Execute(&builder, ctx)
	return builder.String()
}
