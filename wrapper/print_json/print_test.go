package print_json_test

import (
	"encoding/json"
	"testing"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/DmitryBogomolov/dicecalc/wrapper/print_json"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	probs, _ := probabilities.NewProbabilities(2, 5, 7, []uint64{1, 2, 3, 1})
	ret := print_json.Print(probs, "Hello World")
	var obj map[string]any
	assert.NoError(t, json.Unmarshal([]byte(ret), &obj))
	assert.Equal(t, "Hello World", obj["title"])
	values := obj["values"].([]any)
	assert.Equal(t, 4, len(values))
	assert.Equal(t, 2.0, values[0].(map[string]any)["value"])
	assert.Equal(t, 1.0, values[0].(map[string]any)["count"])
	assert.Equal(t, "14.3%", values[0].(map[string]any)["probability"])
	assert.Equal(t, 3.0, values[1].(map[string]any)["value"])
	assert.Equal(t, 2.0, values[1].(map[string]any)["count"])
	assert.Equal(t, "28.6%", values[1].(map[string]any)["probability"])
	assert.Equal(t, 4.0, values[2].(map[string]any)["value"])
	assert.Equal(t, 3.0, values[2].(map[string]any)["count"])
	assert.Equal(t, "42.9%", values[2].(map[string]any)["probability"])
	assert.Equal(t, 5.0, values[3].(map[string]any)["value"])
	assert.Equal(t, 1.0, values[3].(map[string]any)["count"])
	assert.Equal(t, "14.3%", values[3].(map[string]any)["probability"])
}
