package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/DmitryBogomolov/dicecalc/wrapper"
)

const PORT = 3001
const DEFAULT_MODE = "sum"
const DEFAULT_OUTPUT = "html"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/calculate", handleCalculate)
	err := http.ListenAndServe(fmt.Sprintf(":%d", getPort()), mux)
	fmt.Println(err)
}

func getPort() int {
	if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		return port
	}
	return PORT
}

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
	data.Init.DiceCount = 2
	data.Init.DiceSides = 6
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
