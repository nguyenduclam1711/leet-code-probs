package decodeways

import "strconv"

func NumDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1

	for i := len(s) - 1; i >= 0; i-- {
		res := 0
		firstNum, _ := strconv.Atoi(s[i : i+1])

		if firstNum > 0 {
			res += dp[i+1]

			if i <= len(s)-2 {
				secondNum, _ := strconv.Atoi(s[i : i+2])

				if secondNum <= 26 {
					res += dp[i+2]
				}
			}
		}
		dp[i] = res
	}
	return dp[0]
}
