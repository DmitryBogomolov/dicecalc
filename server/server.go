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
	Title     string
	Modes     []string
	Outputs   []string
	DiceCount int
	DiceSides int
}

func handleRoot(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var data _TemplateData
	data.Title = "Test"
	data.Modes = wrapper.Modes()
	data.Outputs = wrapper.Outputs()
	data.DiceCount = 2
	data.DiceSides = 6
	tmpl.Execute(writer, data)
}

var contentTypes = map[string]string{
	"raw":  "text/plain",
	"html": "text/html",
	"json": "application/json",
	"svg":  "image/svg+xml",
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
