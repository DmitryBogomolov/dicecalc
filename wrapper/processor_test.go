package wrapper_test

import (
	"testing"

	"github.com/DmitryBogomolov/dicecalc/wrapper"
	"github.com/stretchr/testify/assert"
)

func TestModes(t *testing.T) {
	assert.Equal(t, []string{"max", "min", "sum"}, wrapper.Modes())
}

func TestOutputs(t *testing.T) {
	assert.Equal(t, []string{"html", "json", "raw", "svg"}, wrapper.Outputs())
}

func TestProcess(t *testing.T) {
	ret, err := wrapper.Process("sum", "1d2", "raw")
	assert.NoError(t, err)
	assert.Equal(t, "Probabilities of 1d2 (sum) rolls\n 1   1  50.0000%\n 2   1  50.0000%\n", string(ret))
}
