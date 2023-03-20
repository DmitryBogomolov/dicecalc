package pages

import (
	_ "embed"
	"io"
	"text/template"

	"github.com/DmitryBogomolov/dicecalc/wrapper"
)

const (
	DEFAULT_DICE_COUNT = 2
	DEFAULT_DICE_SIDES = 6
	DEFAULT_MODE       = "sum"
	DEFAULT_OUTPUT     = "html"
)

//go:embed template.html
var tmplStr string
var tmpl = template.Must(template.New("/").Parse(tmplStr))

type _TemplateData struct {
	Title   string
	Modes   []string
	Outputs []string
	Init    struct {
		DiceCount int
		DiceSides int
		Mode      int
		Output    int
	}
}

func RenderSelection(writer io.Writer) error {
	// TODO: Temporary for debug.
	// fileObj, _ := os.Open("./server/pages/template.html")
	// defer fileObj.Close()
	// tmplStr, _ := ioutil.ReadAll(fileObj)
	// tmpl := template.Must(template.New(".").Parse(string(tmplStr)))

	modes := wrapper.Modes()
	outputs := wrapper.Outputs()

	var data _TemplateData
	data.Title = "Dice roll probabilities calculator"
	data.Modes = modes
	data.Outputs = outputs
	data.Init.DiceCount = DEFAULT_DICE_COUNT
	data.Init.DiceSides = DEFAULT_DICE_SIDES
	data.Init.Mode = findIndex(modes, DEFAULT_MODE)
	data.Init.Output = findIndex(outputs, DEFAULT_OUTPUT)
	return tmpl.Execute(writer, data)
}

func findIndex[T comparable](items []T, target T) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

var contentTypes = map[string]string{
	"raw":  "text/plain",
	"html": "text/html",
	"json": "application/json",
	"svg":  "image/svg+xml",
	"csv":  "text/plain",
}

func GetContentType(output string) string {
	return contentTypes[output]
}

func RenderCalculation(writer io.Writer, schema, mode, output string) error {
	if data, err := wrapper.Process(mode, schema, output); err != nil {
		return err
	} else {
		writer.Write(data)
		return nil
	}
}
