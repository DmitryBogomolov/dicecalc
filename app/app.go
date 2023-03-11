package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DmitryBogomolov/dicecalc/wrapper"
)

func main() {
	modeVar := flag.String("mode", "", "operation")
	schemaVar := flag.String("schema", "", "roll schema")
	outputVar := flag.String("output", "raw", "output format")
	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
		return
	}

	if ret, err := wrapper.Process(*modeVar, *schemaVar, *outputVar); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(ret))
	}
}
