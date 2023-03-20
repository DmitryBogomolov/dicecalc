package main

import (
	"context"
	"fmt"
)

func Handle(ctx context.Context, request []byte) ([]byte, error) {
	fmt.Printf("name    : %s\n", ctx.Value("lambdaRuntimeFunctionName").(string))
	fmt.Printf("version : %s\n", ctx.Value("lambdaRuntimeFunctionVersion").(string))
	fmt.Printf("limit   : %dMb\n", ctx.Value("lambdaRuntimeMemoryLimit").(int))
	fmt.Printf("request : %s\n", ctx.Value("lambdaRuntimeRequestID").(string))
	fmt.Printf("byte: %v\n", request)
	fmt.Printf("str : %v\n", string(request))

	respone := "Hello World\n"

	return []byte(respone), nil
}
