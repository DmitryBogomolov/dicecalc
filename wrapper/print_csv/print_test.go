package print_csv_test

import (
	"testing"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/wrapper/print_csv"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	probs, _ := probabilities.NewProbabilities(2, 5, 7, []uint64{1, 2, 3, 1})
	ret := print_csv.Print(probs, "Hello World")
	assert.Equal(
		t,
		[]byte("# Hello World\n2,1,14.3%\n3,2,28.6%\n4,3,42.9%\n5,1,14.3%\n# Total count: 7\n"),
		ret,
	)
}
