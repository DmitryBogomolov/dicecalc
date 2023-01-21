package dicecalc

type _DiceRoll struct {
	dices []byte
}

func initDiceRoll(value int, params DiceRollParameters) *_DiceRoll {
	dices := make([]byte, params.DiceCount)
	for i := range dices {
		dices[i] = 1
	}
	rest := byte(value - len(dices))
	k := len(dices) - 1
	for rest > 0 {
		val := byte(params.DiceSides - 1)
		if rest < val {
			val = rest
		}
		rest -= val
		dices[k] += val
		k--
	}
	return &_DiceRoll{dices: dices}
}

func (diceRoll *_DiceRoll) key() string {
	return string(diceRoll.dices)
}

func (diceRoll *_DiceRoll) getSimilarRoll(srcIdx, dstIdx int) *_DiceRoll {
	dices := diceRoll.dices
	if dstIdx == srcIdx-1 {
		if dices[srcIdx]-dices[dstIdx] < 2 {
			return nil
		}
	} else {
		if dices[srcIdx]-dices[srcIdx-1] < 1 {
			return nil
		}
		if dices[dstIdx+1]-dices[dstIdx] < 1 {
			return nil
		}
	}
	dices = append([]byte(nil), dices...)
	dices[srcIdx]--
	dices[dstIdx]++
	return &_DiceRoll{dices: dices}
}

func (diceRoll *_DiceRoll) getAllSimilarRolls() []*_DiceRoll {
	var rolls []*_DiceRoll
	for i := len(diceRoll.dices) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if roll := diceRoll.getSimilarRoll(i, j); roll != nil {
				rolls = append(rolls, roll)
			}
		}
	}
	return rolls
}
