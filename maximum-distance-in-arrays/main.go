package maximumdistanceinarrays

import (
	"math"
)

func maxDistance(arrays [][]int) int {
	firstMax, secondMax := math.MinInt, math.MinInt
	firstMaxIndex := -1
	res := 0

	for i, arr := range arrays {
		currMax := arr[len(arr)-1]
		if currMax > firstMax {
			secondMax = firstMax
			firstMax = currMax
			firstMaxIndex = i
		} else if currMax > secondMax {
			secondMax = currMax
		}
	}
	for i, arr := range arrays {
		currMax := firstMax
		if i == firstMaxIndex {
			currMax = secondMax
		}
		newRes := currMax - arr[0]
		if newRes > res {
			res = newRes
		}
	}
	return res
}
