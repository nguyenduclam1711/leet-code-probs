package minimumflipstomakeaorbequaltoc

import (
	"strconv"
)

func addZeroBits(s *string, difLen int) {
	newBit := make([]byte, difLen)
	for i := range newBit {
		newBit[i] = '0'
	}
	*s = string(newBit) + *s
}

func minFlips(a int, b int, c int) int {
	bitA, bitB, bitC := strconv.FormatInt(int64(a), 2), strconv.FormatInt(int64(b), 2), strconv.FormatInt(int64(c), 2)
	maxLen := len(bitA)

	if len(bitA) < len(bitB) {
		addZeroBits(&bitA, len(bitB)-len(bitA))
		maxLen = len(bitB)
	} else if len(bitB) < len(bitA) {
		addZeroBits(&bitB, len(bitA)-len(bitB))
	}

	if maxLen > len(bitC) {
		addZeroBits(&bitC, maxLen-len(bitC))
	} else if len(bitC) > maxLen {
		difLen := len(bitC) - maxLen
		addZeroBits(&bitA, difLen)
		addZeroBits(&bitB, difLen)
		maxLen = len(bitC)
	}

	res := 0
	for i := 0; i < maxLen; i++ {
		if bitC[i] == '1' && bitA[i] == bitB[i] && bitA[i] == '0' {
			res++
		} else if bitC[i] == '0' {
			if bitA[i] != bitB[i] {
				res++
			} else if bitA[i] == bitB[i] && bitA[i] == '1' {
				res += 2
			}
		}
	}
	return res
}
