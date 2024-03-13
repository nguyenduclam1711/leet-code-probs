package mincostclimbingstairs

func findMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)

	for i := 2; i < len(dp); i++ {
		dp[i] = findMin(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(dp)-1]
}
