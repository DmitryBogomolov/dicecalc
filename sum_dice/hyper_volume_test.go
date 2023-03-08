package sum_dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHyperVolume(t *testing.T) {
	assert.Equal(t, 0.0, getHyperVolume(0, 0))

	assert.Equal(t, 0.0, getHyperVolume(0, 1))
	assert.Equal(t, 0.4, getHyperVolume(0.4, 1))
	assert.Equal(t, 0.6, getHyperVolume(0.6, 1))
	assert.Equal(t, 1.0, getHyperVolume(1, 1))

	assert.Equal(t, 0.0, getHyperVolume(0, 2))
	assert.Equal(t, 0.125, getHyperVolume(0.5, 2))
	assert.Equal(t, 0.5, getHyperVolume(1, 2))
	assert.Equal(t, 0.875, getHyperVolume(1.5, 2))
	assert.Equal(t, 1.0, getHyperVolume(2, 2))

	assert.Equal(t, 0.0, getHyperVolume(0, 3))
	assert.Equal(t, 3.0/18.0, getHyperVolume(1, 3))
	assert.Equal(t, 0.5, getHyperVolume(1.5, 3))
	assert.Equal(t, 15.0/18.0, getHyperVolume(2, 3))
	assert.Equal(t, 1.0, getHyperVolume(3, 3))

	assert.Equal(t, 0.0, getHyperVolume(0, 4))
	assert.Equal(t, 3.0/72.0, getHyperVolume(1, 4))
	assert.Equal(t, 0.5, getHyperVolume(2, 4))
	assert.Equal(t, 69.0/72.0, getHyperVolume(3, 4))
	assert.Equal(t, 1.0, getHyperVolume(4, 4))
}
