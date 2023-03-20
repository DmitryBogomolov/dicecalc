package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/server/pages"
	"github.com/DmitryBogomolov/dicecalc/sum_dice"
)

type Request struct {
	HttpMethod                      string              `json:"httpMethod"`
	Headers                         map[string]string   `json:"headers"`
	MultiValueHeaders               map[string][]string `json:"mutliValueHeaders"`
	QueryStringParameters           map[string]string   `json:"queryStringParameters"`
	MultiValueQueryStringParameters map[string][]string `json:"multiValueQueryStringParameters"`
	RequestContext                  map[string]any      `json:"requestContent"`
	Body                            string              `json:"body"`
	IsBase64Encoded                 bool                `json:"isBase64Encoded"`
}

type Response struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"mutliValueHeaders"`
	Body              string              `json:"body"`
	IsBase64Encoded   bool                `json:"isBase64Encoded"`
}

const calculationQuery = `.?mode=${mode}&output=${output}&schema=${diceCount}d${diseSides}`

func Handle(ctx context.Context, data []byte) ([]byte, error) {
	fmt.Printf("name    : %s\n", ctx.Value("lambdaRuntimeFunctionName").(string))
	fmt.Printf("version : %s\n", ctx.Value("lambdaRuntimeFunctionVersion").(string))
	fmt.Printf("limit   : %dMb\n", ctx.Value("lambdaRuntimeMemoryLimit").(int))
	fmt.Printf("request : %s\n", ctx.Value("lambdaRuntimeRequestID").(string))
	fmt.Printf("byte: %v\n", data)
	fmt.Printf("str : %v\n", string(data))

	var req Request
	isHttp := json.Unmarshal(data, &req) == nil

	probs, _ := sum_dice.CalculateProbabilities(probabilities.DiceRollParameters{
		DiceCount: 2,
		DiceSides: 6,
	})

	message := fmt.Sprintf("Probs: %d\nHello World\n", probs.TotalVariants())
	var sb strings.Builder
	pages.RenderSelection(&sb, calculationQuery)

	if isHttp {
		var res Response
		res.StatusCode = 200
		res.Body = sb.String()
		res.Headers = map[string]string{
			"content-type": "text/html",
		}
		return json.Marshal(res)
	}

	return []byte(message), nil
}
