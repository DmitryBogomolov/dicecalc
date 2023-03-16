package main

import (
	_ "embed"
	"html/template"
	"net/http"

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

func handleRoot(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO: Temporary for debug.
	// fileObj, _ := os.Open("./server/template.html")
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
	tmpl.Execute(writer, data)
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

func handleCalculate(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	query := req.URL.Query()
	mode := query.Get("mode")
	output := query.Get("output")
	schema := query.Get("schema")
	if data, err := wrapper.Process(mode, schema, output); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	} else {
		contentType := contentTypes[output]
		writer.Header().Set("Content-Type", contentType)
		writer.Write(data)
	}
}
