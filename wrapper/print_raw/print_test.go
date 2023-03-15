package print_raw_test

import (
	"testing"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/wrapper/print_raw"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	probs, _ := probabilities.NewProbabilities(2, 5, 7, []uint64{1, 2, 3, 1})
	ret := print_raw.Print(probs, "Hello World")
	assert.Equal(
		t,
		[]byte("Hello World\n 2   1    14.3%\n 3   2    28.6%\n 4   3    42.9%\n 5   1    14.3%\nTotal count: 7\n"),
		ret,
	)
}
