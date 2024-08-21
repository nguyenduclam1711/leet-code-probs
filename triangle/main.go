package triangle

import "math"

func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	dp[0] = triangle[0]

	for i := 1; i < len(dp); i++ {
		dp[i] = make([]int, i+1)
		for j := range dp[i] {
			if j == 0 {
				dp[i][j] = triangle[i][j] + dp[i-1][0]
			} else if j == i {
				dp[i][j] = triangle[i][j] + dp[i-1][len(dp[i-1])-1]
			} else {
				firstVal := triangle[i][j] + dp[i-1][j-1]
				secondVal := triangle[i][j] + dp[i-1][j]
				if firstVal < secondVal {
					dp[i][j] = firstVal
				} else {
					dp[i][j] = secondVal
				}
			}
		}
	}
	res := math.MaxInt
	for _, val := range dp[len(dp)-1] {
		if val < res {
			res = val
		}
	}
	return res
}
