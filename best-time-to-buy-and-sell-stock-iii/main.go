package besttimetobuyandsellstockiii

func maxProfit(prices []int) int {
	firstDp, secondDp := make([]int, len(prices)), make([]int, len(prices))

	// calculate first dp
	low := prices[0]
	for i := 1; i < len(firstDp); i++ {
		lastHighProfit := firstDp[i-1]
		currProfit := prices[i] - low
		if currProfit > lastHighProfit {
			firstDp[i] = currProfit
		} else {
			firstDp[i] = lastHighProfit
		}
		if prices[i] < low {
			low = prices[i]
		}
	}

	// calculate second dp
	high := prices[len(prices)-1]
	for i := len(secondDp) - 2; i >= 0; i-- {
		lastHightProfit := secondDp[i+1]
		currProfit := high - prices[i]
		if currProfit > lastHightProfit {
			secondDp[i] = currProfit
		} else {
			secondDp[i] = lastHightProfit
		}
		if prices[i] > high {
			high = prices[i]
		}
	}

	res := 0
	for i := range firstDp {
		var currRes int
		if i < len(firstDp)-1 {
			currRes = firstDp[i] + secondDp[i+1]
		} else {
			currRes = firstDp[i]
		}
		if currRes > res {
			res = currRes
		}
	}
	return res
}
