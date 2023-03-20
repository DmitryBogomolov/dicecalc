package main

import (
	_ "embed"
	"net/http"

	"github.com/DmitryBogomolov/dicecalc/server/pages"
)

const calculationQuery = `./calculate?mode=${mode}&output=${output}&schema=${diceCount}d${diseSides}`

func handleRoot(writer http.ResponseWriter, req *http.Request) {
	if !checkMethod(writer, req) {
		return
	}
	if err := pages.RenderSelection(writer, calculationQuery); err != nil {
		sendError(writer, err)
	}
}

func handleCalculate(writer http.ResponseWriter, req *http.Request) {
	if !checkMethod(writer, req) {
		return
	}
	query := req.URL.Query()
	mode := query.Get("mode")
	output := query.Get("output")
	schema := query.Get("schema")
	contentType := pages.GetContentType(output)
	writer.Header().Set("Content-Type", contentType)
	if err := pages.RenderCalculation(writer, schema, mode, output); err != nil {
		sendError(writer, err)
	}
}

func checkMethod(writer http.ResponseWriter, req *http.Request) bool {
	isValid := req.Method == http.MethodGet
	if !isValid {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
	return isValid
}

func sendError(writer http.ResponseWriter, err error) {
	http.Error(writer, err.Error(), http.StatusBadRequest)
}
