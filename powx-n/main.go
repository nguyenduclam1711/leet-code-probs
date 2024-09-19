package powxn

import "math"

func myPow(x float64, n int) float64 {
	res := 1.0
	step := math.Abs(float64(n))
	for i := 0; i < int(step); i++ {
		res *= x
	}
	if n < 0 {
		return 1 / res
	}
	return res
}
