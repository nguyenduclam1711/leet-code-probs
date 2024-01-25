package stringcompression

import (
	"math"
	"strconv"
)

func appendChars(currChar byte, insertPos *int, chars []byte, count int) {
	chars[*insertPos] = currChar
	*insertPos++

	if count > 1 {
		// get exponent of 10
		exponent := 0
		for math.Pow10(exponent) <= float64(count) {
			exponent++
		}
		// to get the curr exponent
		exponent--

		// append the count numbers
		for j := exponent; j >= 0; j-- {
			pow10 := int(math.Pow10(j))
			appendNumber := count / pow10
			chars[*insertPos] = strconv.Itoa(appendNumber)[0]
			*insertPos++
			count = count - appendNumber*pow10
		}
	}
}

func compress(chars []byte) int {
	insertPos := 0
	currChar := chars[0]
	count := 1

	for i := 1; i < len(chars); i++ {
		if chars[i] != currChar {
			appendChars(currChar, &insertPos, chars, count)
			count = 1
			currChar = chars[i]
		} else {
			count++
		}
	}
	appendChars(currChar, &insertPos, chars, count)
	return insertPos
}
