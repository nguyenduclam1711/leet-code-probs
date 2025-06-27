package besttimetobuyandsellstockii

func maxProfit(prices []int) int {
	// to get all the profit by checking the last zero price (cover the edge case: the last element is larger than the previous pointer price)
	prices = append(prices, 0)
	res := 0
	i, j := 0, 1
	currMax := prices[1]
	for i < len(prices) && j < len(prices) {
		if prices[i] >= prices[j] {
			if currMax > prices[i] {
				res += currMax - prices[i]
			}
			i = j
			j = i + 1
			if j < len(prices) {
				currMax = prices[j]
			}
		} else {
			if prices[j] > currMax {
				currMax = prices[j]
				j++
			} else {
				res += currMax - prices[i]
				i = j
				j = i + 1
				if j < len(prices) {
					currMax = prices[j]
				}
			}
		}
	}
	return res
}
