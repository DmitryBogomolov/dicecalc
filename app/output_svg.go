package main

import (
	"html/template"
	"strings"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
)

const tmplStr = `<svg width="{{ .Width }}" height="{{ .Height }}" xmlns="http://www.w3.org/2000/svg">
  <text x="20" y="40">{{ .Title }}</text>
  <path class="arg-axis" stroke="red" stroke-width="1" d="{{ .ArgAxisPath }}" />
  <path class="val-axis" stroke="red" stroke-width="1" d="{{ .ValAxisPath }}" />
  <g class="arg-ticks">
  {{ range .ArgTicks }}
	<path  />
  {{ end }}
  </g>
  <g class="var-ticks">
  {{ range .ValTicks }}
    <path />
  {{ end }}
  </g>
  <g class="arg-labels">
  {{ range .ArgLabels }}
    <text>{{ .Text }}</text>
  {{ end }}
  </g>
  <g class="val-labels">
  {{ range .ValLabels }}
    <text>{{ .Text }}</text>
  {{ end }}
  </g>
  <path class="data" stroke="blue" stroke-width="2" d="{{ .DataPath }}" />
</svg>`

var tmpl = template.Must(template.New("svg").Parse(tmplStr))

type _TmplContent struct {
	Width       int
	Height      int
	Title       string
	ArgAxisPath string
	ValAxisPath string
	ArgLabels   []_Label
	ValLabels   []_Label
	ArgTicks    []_Tick
	ValTicks    []_Tick
	DataPath    string
}

type _Tick struct{}
type _Label struct {
	Text string
}

func displaySvg(probs probabilities.Probabilities, title string) string {
	var builder strings.Builder
	var ctx _TmplContent
	ctx.Width = 640
	ctx.Height = 480
	ctx.Title = title
	ctx.ArgAxisPath = "M 50 450 L 600 450"
	ctx.ValAxisPath = "M 50 50 L 50 450"
	ctx.ArgTicks = []_Tick{
		{},
		{},
	}
	ctx.ValTicks = []_Tick{
		{},
		{},
		{},
	}
	ctx.ArgLabels = []_Label{
		{Text: "1"},
		{Text: "2"},
	}
	ctx.ValLabels = []_Label{
		{Text: "A"},
		{Text: "B"},
		{Text: "C"},
	}
	ctx.DataPath = "M 0 0 L 100 100 Z"
	tmpl.Execute(&builder, ctx)
	return builder.String()
}
