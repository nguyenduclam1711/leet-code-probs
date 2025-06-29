package besttimetobuyandsellstockiii

func maxProfit(prices []int) int {
	firstDp, secondDp := make([]int, len(prices)), make([]int, len(prices))
	low := prices[0]
	high := prices[len(prices)-1]
	for i := 1; i < len(prices); i++ {
		firstDp[i] = max(prices[i]-low, firstDp[i-1])
		if prices[i] < low {
			low = prices[i]
		}
		j := len(prices) - i - 1
		secondDp[j] = max(high-prices[j], secondDp[j+1])
		if prices[j] > high {
			high = prices[j]
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
