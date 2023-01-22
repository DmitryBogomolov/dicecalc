package dicecalc_test

import (
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/DmitryBogomolov/dicecalc"
)

func makeProbabilitiesList(probabilities *Probabilities) []float64 {
	probs := make([]float64, probabilities.ValuesCount())
	for i := range probs {
		probs[i] = probabilities.ValueProbability(i + probabilities.MinValue())
	}
	return probs
}

func checkProbabilities(t *testing.T, name string, expected []float64) {
	t.Run(name, func(t *testing.T) {
		parts := strings.Split(name, "d")
		count, _ := strconv.Atoi(parts[0])
		sides, _ := strconv.Atoi(parts[1])
		probabilities, _ := CalculateProbabilities(DiceRollParameters{DiceCount: count, DiceSides: sides})
		actual := makeProbabilitiesList(probabilities)
		for i := range expected {
			if math.Abs(actual[i]-expected[i]) > 0.0001 {
				t.Errorf("%dth element: expected %f actual %f", i, expected[i], actual[i])
			}
		}
	})
}

func TestCalculateProbabilities(t *testing.T) {
	probabilities, _ := CalculateProbabilities(DiceRollParameters{DiceCount: 3, DiceSides: 6})
	assert.Equal(t, 3, probabilities.MinValue())
	assert.Equal(t, 18, probabilities.MaxValue())
	assert.Equal(t, 16, probabilities.ValuesCount())
}

func TestValidation(t *testing.T) {
	{
		probabilities, err := CalculateProbabilities(DiceRollParameters{DiceCount: 0, DiceSides: 4})
		assert.Nil(t, probabilities)
		assert.Error(t, err)
	}
	{
		probabilities, err := CalculateProbabilities(DiceRollParameters{DiceCount: 3, DiceSides: 0})
		assert.Nil(t, probabilities)
		assert.Error(t, err)
	}
	{
		probabilities, err := CalculateProbabilities(DiceRollParameters{DiceCount: MAX_DICE_COUNT + 1, DiceSides: 4})
		assert.Nil(t, probabilities)
		assert.Error(t, err)
	}
	{
		probabilities, err := CalculateProbabilities(DiceRollParameters{DiceCount: 3, DiceSides: MAX_DICE_SIDES + 1})
		assert.Nil(t, probabilities)
		assert.Error(t, err)
	}

}

func Test3dx(t *testing.T) {
	checkProbabilities(t, "3d4", []float64{
		0.0156,
		0.0469,
		0.0938,
		0.1562,
		0.1875,
		0.1875,
		0.1562,
		0.0938,
		0.0469,
		0.0156,
	})
	checkProbabilities(t, "3d6", []float64{
		0.0046,
		0.0139,
		0.0278,
		0.0463,
		0.0694,
		0.0972,
		0.1157,
		0.1250,
		0.1250,
		0.1157,
		0.0972,
		0.0694,
		0.0463,
		0.0278,
		0.0139,
		0.0046,
	})
	checkProbabilities(t, "3d10", []float64{
		0.0010,
		0.0030,
		0.0060,
		0.0100,
		0.0150,
		0.0210,
		0.0280,
		0.0360,
		0.0450,
		0.0550,
		0.0630,
		0.0690,
		0.0730,
		0.0750,
		0.0750,
		0.0730,
		0.0690,
		0.0630,
		0.0550,
		0.0450,
		0.0360,
		0.0280,
		0.0210,
		0.0150,
		0.0100,
		0.0060,
		0.0030,
		0.0010,
	})
	checkProbabilities(t, "3d20", []float64{
		0.0001,
		0.0004,
		0.0008,
		0.0013,
		0.0019,
		0.0026,
		0.0035,
		0.0045,
		0.0056,
		0.0069,
		0.0083,
		0.0097,
		0.0114,
		0.0131,
		0.0150,
		0.0170,
		0.0191,
		0.0214,
		0.0238,
		0.0262,
		0.0285,
		0.0305,
		0.0323,
		0.0338,
		0.0350,
		0.0360,
		0.0367,
		0.0372,
		0.0375,
		0.0375,
		0.0372,
		0.0367,
		0.0360,
		0.0350,
		0.0338,
		0.0323,
		0.0305,
		0.0285,
		0.0262,
		0.0238,
		0.0214,
		0.0191,
		0.0170,
		0.0150,
		0.0131,
		0.0114,
		0.0097,
		0.0083,
		0.0069,
		0.0056,
		0.0045,
		0.0035,
		0.0026,
		0.0019,
		0.0013,
		0.0008,
		0.0004,
		0.0001,
	})
}

func Test4dx(t *testing.T) {
	checkProbabilities(t, "4d4", []float64{
		0.0039,
		0.0156,
		0.0391,
		0.0781,
		0.1211,
		0.1562,
		0.1719,
		0.1562,
		0.1211,
		0.0781,
		0.0391,
		0.0156,
		0.0039,
	})
	checkProbabilities(t, "4d6", []float64{
		0.0008,
		0.0031,
		0.0077,
		0.0154,
		0.0270,
		0.0432,
		0.0617,
		0.0802,
		0.0965,
		0.1080,
		0.1127,
		0.1080,
		0.0965,
		0.0802,
		0.0617,
		0.0432,
		0.0270,
		0.0154,
		0.0077,
		0.0031,
		0.0008,
	})
	checkProbabilities(t, "4d10", []float64{
		0.0001,
		0.0004,
		0.0010,
		0.0020,
		0.0035,
		0.0056,
		0.0084,
		0.0120,
		0.0165,
		0.0220,
		0.0282,
		0.0348,
		0.0415,
		0.0480,
		0.0540,
		0.0592,
		0.0633,
		0.0660,
		0.0670,
		0.0660,
		0.0633,
		0.0592,
		0.0540,
		0.0480,
		0.0415,
		0.0348,
		0.0282,
		0.0220,
		0.0165,
		0.0120,
		0.0084,
		0.0056,
		0.0035,
		0.0020,
		0.0010,
		0.0004,
		0.0001,
	})
}
