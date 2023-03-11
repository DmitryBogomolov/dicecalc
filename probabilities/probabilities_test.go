package probabilities_test

import (
	"testing"

	"github.com/DmitryBogomolov/dicecalc/probabilities"
	"github.com/stretchr/testify/assert"
)

func TestProbabilities(t *testing.T) {
	probs, _ := probabilities.NewProbabilities(2, 5, uint64(10), []uint64{1, 2, 4, 3})
	assert.Equal(t, 2, probs.MinValue())
	assert.Equal(t, 5, probs.MaxValue())
	assert.Equal(t, 0.1, probs.MinProbability())
	assert.Equal(t, 0.4, probs.MaxProbability())
	assert.Equal(t, 4, probs.Count())
	assert.Equal(t, uint64(10), probs.TotalVariants())

	assert.Equal(t, uint64(1), probs.ValueVariants(2))
	assert.Equal(t, uint64(2), probs.ValueVariants(3))
	assert.Equal(t, uint64(4), probs.ValueVariants(4))
	assert.Equal(t, uint64(3), probs.ValueVariants(5))

	assert.Equal(t, uint64(0), probs.ValueVariants(0))
	assert.Equal(t, uint64(0), probs.ValueVariants(7))

	assert.Equal(t, 0.1, probs.ValueProbability(2))
	assert.Equal(t, 0.2, probs.ValueProbability(3))
	assert.Equal(t, 0.4, probs.ValueProbability(4))
	assert.Equal(t, 0.3, probs.ValueProbability(5))

	assert.Equal(t, 0.0, probs.ValueProbability(0))
	assert.Equal(t, 0.0, probs.ValueProbability(7))
}

func TestProbabilitiesChecks(t *testing.T) {
	var err error

	_, err = probabilities.NewProbabilities(2, 1, uint64(1), nil)
	assert.ErrorContains(t, err, "bad value range 2..1")

	_, err = probabilities.NewProbabilities(2, 5, uint64(1), nil)
	assert.ErrorContains(t, err, "bad variants - length should be 4")

	_, err = probabilities.NewProbabilities(2, 5, uint64(2), []uint64{1, 2, 3, 4})
	assert.ErrorContains(t, err, "bad total 2 - should be 10")
}
