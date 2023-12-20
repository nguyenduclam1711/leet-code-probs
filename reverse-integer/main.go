package reverseinteger

import (
	"fmt"
	"math"
	"strconv"
)

func Reverse(x int) int {
	max := int(math.Pow(2, 31))
	isLessThanZero := x < 0
	if isLessThanZero {
		x = -x
	}
	str := fmt.Sprintf("%d", x)
	res := 0

	for i := len(str) - 1; i >= 0; i-- {
		if num, err := strconv.Atoi(string(str[i])); err == nil {
			res += num * int(math.Pow(10, float64(i)))
		}
		if res >= max {
			return 0
		}
	}
	if isLessThanZero {
		return -res
	}
	return res
}
