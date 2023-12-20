package reversebits

import (
	"math"
)

func ReverseBits(num uint32) uint32 {
	var result uint32 = 0
	pos := 31

	for num > 0 {
		result += num & 1 * uint32(math.Pow(2, float64(pos)))
		num >>= 1
		pos--
	}
	return result
}
