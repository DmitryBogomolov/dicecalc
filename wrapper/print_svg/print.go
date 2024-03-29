package print_svg

import (
	_ "embed"
	"fmt"
	"html/template"
	"math"
	"strings"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/wrapper/util"
)

const (
	WIDTH             = 640
	HEIGHT            = 480
	PADDING           = 4
	BORDER_COLOR      = "#000"
	TITLE_COLOR       = "#000"
	TITLE_SIZE        = 24
	TITLE_PADDING     = 8
	AXIS_COLOR        = "#222"
	AXIS_PADDING      = 16
	TICK_SIZE         = 8
	ARG_TICK_COUNT    = 5
	VAL_TICK_COUNT    = 5
	LABEL_COLOR       = "#000"
	LABEL_OFFSET      = 4
	LABEL_SIZE        = 16
	ARG_LABEL_PADDING = LABEL_SIZE
	VAL_LABEL_PADDING = 2 * LABEL_SIZE
	POINT_COLOR       = "#00F"
	POINT_SIZE        = 4
)

//go:embed template.svg
var tmplStr string

var tmpl = template.Must(template.New("svg").Parse(tmplStr))

type _TemplateData struct {
	Size        _Point
	BorderColor string
	Title       _Title
	ArgAxis     _Axis
	ValAxis     _Axis
	DataItems   []_DataItem
}

type _Point struct {
	X int
	Y int
}
type _Title struct {
	Position _Point
	Text     string
	Size     int
	Color    string
}
type _Axis struct {
	Color  string
	Path   string
	Ticks  []_Tick
	Labels []_Label
}
type _Tick struct {
	Path  string
	Color string
}
type _Label struct {
	_Point
	Text  string
	Color string
	Size  int
}
type _DataItem struct {
	_Point
	Color       string
	Size        int
	Text        string
	Value       int
	Probability float64
}

func Print(probs probabilities.Probabilities, title string) []byte {
	var builder strings.Builder
	data := makeTemplateData(probs, title)
	tmpl.Execute(&builder, data)
	return []byte(builder.String())
}

func makeTemplateData(probs probabilities.Probabilities, title string) _TemplateData {
	totalX := WIDTH
	totalY := HEIGHT
	minX := PADDING
	maxX := totalX - PADDING
	minY := PADDING + TITLE_SIZE + TITLE_PADDING
	maxY := totalY - PADDING

	dataX1 := minX + LABEL_SIZE + LABEL_OFFSET + AXIS_PADDING + VAL_LABEL_PADDING
	dataX2 := maxX - ARG_LABEL_PADDING
	dataY1 := minY
	dataY2 := maxY - LABEL_SIZE - LABEL_OFFSET - AXIS_PADDING

	argAxisX1 := dataX1
	argAxixX2 := dataX2
	argAxixY := dataY2 + AXIS_PADDING
	valAxisX := dataX1 - AXIS_PADDING
	valAxisY1 := dataY1
	valAxisY2 := dataY2

	var data _TemplateData
	data.Size = _Point{totalX, totalY}
	data.BorderColor = BORDER_COLOR
	data.Title = _Title{
		Position: _Point{
			X: (minX + maxX) / 2,
			Y: PADDING + (TITLE_SIZE / 2),
		},
		Text:  title,
		Size:  TITLE_SIZE,
		Color: TITLE_COLOR,
	}
	data.ArgAxis = _Axis{
		Color:  AXIS_COLOR,
		Path:   fmt.Sprintf("M %d %d L %d %d", argAxisX1, argAxixY, argAxixX2, argAxixY),
		Ticks:  collectArgTicks(argAxisX1, argAxixX2, argAxixY, ARG_TICK_COUNT),
		Labels: collectArgLabels(probs, argAxisX1, argAxixX2, argAxixY, ARG_TICK_COUNT),
	}
	data.ValAxis = _Axis{
		Color:  AXIS_COLOR,
		Path:   fmt.Sprintf("M %d %d L %d %d", valAxisX, valAxisY1, valAxisX, valAxisY2),
		Ticks:  collectValTicks(valAxisX, valAxisY1, valAxisY2, VAL_TICK_COUNT),
		Labels: collectValLabels(probs, valAxisX, valAxisY1, valAxisY2, VAL_TICK_COUNT),
	}
	data.DataItems = collectData(probs, dataX1, dataX2, dataY1, dataY2)

	return data
}

func flerp(a, b, k float64) float64 {
	return a*(1-k) + b*k
}

func ilerp(a, b int, k float64) int {
	c := math.Round(flerp(float64(a), float64(b), k))
	return int(c)
}

func collectArgTicks(x1, x2, y, count int) []_Tick {
	ticks := make([]_Tick, count)
	for i := 0; i < count; i++ {
		k := float64(i) / float64(count-1)
		x := ilerp(x1, x2, k)
		ticks[i] = _Tick{
			Path:  fmt.Sprintf("M %d %d L %d %d", x, y, x, y-TICK_SIZE),
			Color: AXIS_COLOR,
		}
	}
	return ticks
}

func collectValTicks(x, y1, y2, count int) []_Tick {
	ticks := make([]_Tick, count)
	for i := 0; i < count; i++ {
		k := float64(i) / float64(count-1)
		y := ilerp(y2, y1, k)
		ticks[i] = _Tick{
			Path:  fmt.Sprintf("M %d %d L %d %d", x, y, x+TICK_SIZE, y),
			Color: AXIS_COLOR,
		}
	}
	return ticks
}

func collectArgLabels(probs probabilities.Probabilities, x1, x2, y, count int) []_Label {
	labels := make([]_Label, count)
	for i := 0; i < count; i++ {
		k := float64(i) / float64(count-1)
		val := ilerp(probs.MinValue(), probs.MaxValue(), k)
		labels[i] = _Label{
			_Point: _Point{
				ilerp(x1, x2, k),
				y + LABEL_OFFSET,
			},
			Text:  fmt.Sprintf("%d", val),
			Color: LABEL_COLOR,
			Size:  LABEL_SIZE,
		}
	}
	return labels
}

func collectValLabels(probs probabilities.Probabilities, x, y1, y2, count int) []_Label {
	labels := make([]_Label, count)
	for i := 0; i < count; i++ {
		k := float64(i) / float64(count-1)
		prob := flerp(probs.MinProbability(), probs.MaxProbability(), k)
		labels[i] = _Label{
			_Point: _Point{
				x - LABEL_OFFSET,
				ilerp(y2, y1, k),
			},
			Text:  fmt.Sprintf("%.2f%%", prob*100),
			Color: LABEL_COLOR,
			Size:  LABEL_SIZE,
		}
	}
	return labels
}

func collectData(probs probabilities.Probabilities, x1, x2, y1, y2 int) []_DataItem {
	minVal := probs.MinValue()
	maxVal := probs.MaxValue()
	minProb := probs.MinProbability()
	maxProb := probs.MaxProbability()
	formatProb := util.GetProbabilityFormatter(probs)
	items := make([]_DataItem, probs.Count())
	for i := 0; i < probs.Count(); i++ {
		val, _, prob := probs.Item(i)
		items[i] = _DataItem{
			_Point: _Point{
				mapValue(float64(val), float64(minVal), float64(maxVal), x1, x2),
				mapValue(prob, minProb, maxProb, y2, y1),
			},
			Value:       val,
			Probability: prob,
			Text:        fmt.Sprintf("%d (%s)", val, formatProb(prob)),
			Color:       POINT_COLOR,
			Size:        POINT_SIZE,
		}
	}
	return items
}

func mapValue(val, minVal, maxVal float64, minTarget, maxTarget int) int {
	k := (val - minVal) / (maxVal - minVal)
	p := k*float64(maxTarget-minTarget) + float64(minTarget)
	return int(math.Round(p))
}
