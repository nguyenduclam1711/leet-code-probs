package flipcolumnsformaximumnumberofequalrows

import "fmt"

func count(s []int, countPattern map[string]int) {
	currN := -1
	currNCount := 1
	patternKey := ""

	for _, n := range s {
		if n != currN {
			patternKey += fmt.Sprint(",", currNCount)
			currN = n
			currNCount = 1
		} else {
			currNCount++
		}
	}
	patternKey += fmt.Sprint(",", currNCount)
	countPattern[patternKey]++
}

func maxEqualRowsAfterFlips(matrix [][]int) int {
	countPattern := map[string]int{}

	for _, r := range matrix {
		count(r, countPattern)
	}
	res := 0
	for _, v := range countPattern {
		if v > res {
			res = v
		}
	}
	return res
}
