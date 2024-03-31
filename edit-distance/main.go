package editdistance

// Can't think of a solution so I used this solution (https://leetcode.com/problems/edit-distance/solutions/3231533/golang-dynamic-programming-with-explanation/?envType=study-plan-v2)
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = findMin(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func findMin(nums ...int) int {
	res := nums[0]
	for _, n := range nums {
		if n < res {
			res = n
		}
	}
	return res
}
