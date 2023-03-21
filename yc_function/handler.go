package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/DmitryBogomolov/dicecalc/server/pages"
)

type Request struct {
	HttpMethod                      string              `json:"httpMethod"`
	Headers                         map[string]string   `json:"headers"`
	QueryStringParameters           map[string]string   `json:"queryStringParameters"`
	MultiValueQueryStringParameters map[string][]string `json:"multiValueQueryStringParameters"`
	RequestContext                  map[string]any      `json:"requestContent"`
	Body                            string              `json:"body"`
	IsBase64Encoded                 bool                `json:"isBase64Encoded"`
	MultiValueHeaders               map[string][]string `json:"mutliValueHeaders"`
}

type Response struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"mutliValueHeaders"`
	Body              string              `json:"body"`
	IsBase64Encoded   bool                `json:"isBase64Encoded"`
}

const calculationQuery = `?mode=${mode}&output=${output}&schema=${diceCount}d${diseSides}`

func Handle(ctx context.Context, req *Request) (*Response, error) {
	// fmt.Printf("name    : %s\n", ctx.Value("lambdaRuntimeFunctionName").(string))
	// fmt.Printf("version : %s\n", ctx.Value("lambdaRuntimeFunctionVersion").(string))
	// fmt.Printf("limit   : %dMb\n", ctx.Value("lambdaRuntimeMemoryLimit").(int))
	// fmt.Printf("request : %s\n", ctx.Value("lambdaRuntimeRequestID").(string))

	schema := req.QueryStringParameters["schema"]
	mode := req.QueryStringParameters["mode"]
	output := req.QueryStringParameters["output"]

	var sb strings.Builder
	var err error
	if schema != "" || mode != "" || output != "" {
		err = pages.RenderCalculation(&sb, schema, mode, output)
	} else {
		funcName := ctx.Value("lambdaRuntimeFunctionName").(string)
		err = pages.RenderSelection(&sb, fmt.Sprintf("/%s%s", funcName, calculationQuery))
	}

	var res Response
	if err != nil {
		res.StatusCode = 500
		res.Body = err.Error()
	} else {
		res.StatusCode = 200
		res.Headers = map[string]string{
			"content-type": "text/html",
		}
		res.Body = sb.String()
	}
	return &res, nil
}
