package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/DmitryBogomolov/dicecalc/wrapper"
)

func main() {
	modes := wrapper.Modes()
	outputs := wrapper.Outputs()

	modeVar := flag.String("mode", "", fmt.Sprintf("mode [%s]", strings.Join(modes, " | ")))
	schemaVar := flag.String("schema", "", "roll schema MdN [1d4, 2d6, 3d8, ...]")
	outputVar := flag.String("output", outputs[0], fmt.Sprintf("output format [%s]", strings.Join(outputs, " | ")))
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
