package maximalsquare

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	edgeLen := 0
	for i := range matrix {
		for j := range matrix[i] {
			squareVal := matrix[i][j]
			if squareVal == '0' {
				dp[i][j] = 0
			} else if i == 0 || j == 0 {
				dp[i][j] = 1
			} else if squareVal == '1' {
				leftVal, upVal, diagonalVal := dp[i][j-1], dp[i-1][j-1], dp[i-1][j]
				minVal := leftVal
				if upVal < minVal {
					minVal = upVal
				}
				if diagonalVal < minVal {
					minVal = diagonalVal
				}
				if minVal == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = minVal + 1
				}
			}
			if dp[i][j] > edgeLen {
				edgeLen = dp[i][j]
			}
		}
	}
	return edgeLen * edgeLen
}
