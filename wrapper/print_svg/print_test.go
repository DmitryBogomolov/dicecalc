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
	assert.Contains(t, string(ret), ">Hello World</text>")
	assert.Equal(t, strings.Count(string(ret), "<circle"), 4)
	assert.Contains(t, string(ret), "<title>2 (14.3%)</title>")
	assert.Contains(t, string(ret), "<title>3 (28.6%)</title>")
	assert.Contains(t, string(ret), "<title>4 (42.9%)</title>")
	assert.Contains(t, string(ret), "<title>5 (14.3%)</title>")
}
