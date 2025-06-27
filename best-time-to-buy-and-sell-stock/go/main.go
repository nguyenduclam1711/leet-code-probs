package besttimetobuyandsellstock

func MaxProfit(prices []int) int {
	profit, currentProfit := 0, 0
	buyDay := 0

	for i := 1; i < len(prices); i++ {
		currentProfit = prices[i] - prices[buyDay]
		if currentProfit > profit {
			profit = currentProfit
		} else if currentProfit < 0 {
			buyDay = i
		}
	}
	return profit
}
