package wrapper_test

import (
	"testing"

	"github.com/DmitryBogomolov/dicecalc/wrapper"
	"github.com/stretchr/testify/assert"
)

func TestModes(t *testing.T) {
	assert.Equal(t, []string{"sum", "min", "max"}, wrapper.Modes())
}

func TestOutputs(t *testing.T) {
	assert.Equal(t, []string{"raw", "json", "html", "svg"}, wrapper.Outputs())
}

func TestProcess(t *testing.T) {
	ret, err := wrapper.Process("sum", "1d2", "raw")
	assert.NoError(t, err)
	assert.Equal(t, "Probabilities of 1d2 (sum) rolls\n 1   1  50.0000%\n 2   1  50.0000%\n", string(ret))
}

func TestProcessErrors(t *testing.T) {
	var err error

	_, err = wrapper.Process("", "", "")
	assert.ErrorContains(t, err, "bad mode: '', bad schema: '', bad output: ''")

	_, err = wrapper.Process("test", "4d10", "raw")
	assert.ErrorContains(t, err, "bad mode: 'test'")
}
