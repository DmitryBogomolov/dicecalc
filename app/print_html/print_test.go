package print_html_test

import (
	"strings"
	"testing"

	"github.com/DmitryBogomolov/dicecalc/app/print_html"
	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	probs, _ := probabilities.NewProbabilities(2, 5, 7, []uint64{1, 2, 3, 1})
	ret := print_html.Print(probs, "Hello World")
	assert.Contains(t, ret, "<title>Hello World</title>")
	assert.Equal(t, 5, strings.Count(ret, "<tr>"))
}
