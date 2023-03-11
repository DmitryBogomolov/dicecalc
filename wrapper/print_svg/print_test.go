package print_svg_test

import (
	"strings"
	"testing"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/wrapper/print_svg"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	probs, _ := probabilities.NewProbabilities(2, 5, 7, []uint64{1, 2, 3, 1})
	ret := print_svg.Print(probs, "Hello World")
	assert.Contains(t, ret, ">Hello World</text>")
	assert.Equal(t, strings.Count(ret, "<circle"), 4)
}
