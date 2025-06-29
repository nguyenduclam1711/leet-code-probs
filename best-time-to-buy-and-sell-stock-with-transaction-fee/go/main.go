package besttimetobuyandsellstockwithtransactionfee

func findMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Can't solve this problem on my own so I have read this solution (https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/solutions/1113471/golang-dp-solution-with-explanation/?envType=study-plan-v2&envId=leetcode-75)
func maxProfit(prices []int, fee int) int {
	dp := make([][2]int, len(prices))
	dp[0][0] = 0          // store the maximum profit when not holding the stock
	dp[0][1] = -prices[0] // store the maximum profit when holding the stock

	for i := 1; i < len(prices); i++ {
		// when not holding the stock, there are 2 options:
		// opt 1: sell the stock held from previous
		// opt 2: still not buy the stock
		dp[i][0] = findMax(prices[i]+dp[i-1][1]-fee, dp[i-1][0])

		// when holding the stock, there are 2 options:
		// opt1: buy stock, based on previous profit
		// opt2: do nothing, keep holding the stock
		dp[i][1] = findMax(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}
