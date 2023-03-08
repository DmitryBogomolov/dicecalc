package sum_dice

import (
	"math"
)

// getHyperVolume returns volume of unit size hyper cube (of dimension n) part
// formed by intersection with `x_1 + ... + x_n = t` plane.
//
// t = 0 - plane intersects cube at (0, ..., 0).
// t = 1 - at (1, ..., 0), ..., (0, ..., 1).
// t = 2 - at (1, 1, ..., 0), ..., (0, ..., 1, 1)
// t = n - at (1, ..., 1)
//
// The formula is
// (1 / n!) * sum_{i=0}^{floor(t)} ( (-1)^i * C_n_i * (t - i)^n )
func getHyperVolume(t float64, n int) float64 {
	if n < 1 {
		return 0
	}
	if t <= 0.0 {
		return 0
	}
	if t >= float64(n) {
		return 1
	}
	m := int(math.Floor(t))
	sum := 0.0
	sign := 1.0
	coeff := 1.0
	for i := 0; i <= m; i++ {
		sum += sign * coeff * math.Pow(t-float64(i), float64(n))
		sign *= -1
		coeff *= float64(n-i) / float64(i+1)
	}
	for i := 2; i <= n; i++ {
		sum /= float64(i)
	}
	return sum
}

// var factorials []uint64 = []uint64{1, 1}

// func getFactorial(k int) uint64 {
// 	if len(factorials) >= (k + 1) {
// 		return factorials[k]
// 	}
// 	ret := uint64(k) * getFactorial((k - 1))
// 	factorials = append(factorials, ret)
// 	return ret
// }