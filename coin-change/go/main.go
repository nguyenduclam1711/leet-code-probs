package coinchange

import "sort"

func findMin(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func CoinChange(coins []int, amount int) int {
	max := amount + 1
	s := make([]int, max)
	sort.Ints(coins)

	for i := range s {
		if i == 0 {
			s[i] = 0
			continue
		}
		// init s[i]
		s[i] = max
		for _, coin := range coins {
			if i == coin {
				s[i] = 1
				break
			}
			if i > coin {
				s[i] = findMin(1+s[i-coin], s[i])
			} else {
				break
			}
		}
	}
	if s[amount] == max {
		return -1
	}
	return s[amount]
}
