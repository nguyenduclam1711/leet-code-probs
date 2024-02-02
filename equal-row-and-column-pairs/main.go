package equalrowandcolumnpairs

import (
	"strconv"
)

func equalPairs(grid [][]int) int {
	rowStr := map[string]int{}
	colStrs := make([]string, len(grid))
	for _, row := range grid {
		str := ""
		for j, n := range row {
			addStr := strconv.Itoa(n) + ","
			str += addStr
			colStrs[j] += addStr
		}
		rowStr[str]++
	}

	res := 0
	for _, str := range colStrs {
		if rowStr[str] > 0 {
			res += rowStr[str]
		}
	}
	return res
}
