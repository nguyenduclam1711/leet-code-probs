package minimumpathsum

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range dp {
		for j := range dp[i] {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				firstVal := grid[i][j] + dp[i][j-1]
				secondVal := grid[i][j] + dp[i-1][j]
				if firstVal < secondVal {
					dp[i][j] = firstVal
				} else {
					dp[i][j] = secondVal
				}
			}
		}
	}
	return dp[m-1][n-1]
}
