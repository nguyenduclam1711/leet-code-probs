package main

import "fmt"

func maxProfit(prices []int) int {
	var profit, currentProfit = 0, 0
	var buyDay = 0

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

func main() {
	fmt.Println("Input: prices = [7,1,5,3,6,4]")
	fmt.Println("Output", maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
