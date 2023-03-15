package print_html_test

import (
	"strings"
	"testing"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/wrapper/print_html"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	probs, _ := probabilities.NewProbabilities(2, 5, 7, []uint64{1, 2, 3, 1})
	ret := print_html.Print(probs, "Hello World")
	assert.Contains(t, string(ret), "<title>Hello World</title>")
	assert.Contains(t, string(ret), "<h2>Hello World</h2>")
	assert.Equal(t, 5, strings.Count(string(ret), "<tr>"))
	assert.Contains(t, string(ret), "<td>2</td><td>1</td><td>14.3%</td>")
	assert.Contains(t, string(ret), "<td>3</td><td>2</td><td>28.6%</td>")
	assert.Contains(t, string(ret), "<td>4</td><td>3</td><td>42.9%</td>")
	assert.Contains(t, string(ret), "<td>5</td><td>1</td><td>14.3%</td>")
	assert.Contains(t, string(ret), "<p>Total count: 7</p>")
}
