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
