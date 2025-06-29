package combinationsum

import (
	"fmt"
	"sort"
)

func makeKey(nums []int) string {
	result := ""
	for i, num := range nums {
		result += fmt.Sprint(num)
		if i < len(nums) {
			result += ","
		}
	}
	return result
}

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	dp := make([][][]int, target+1)

	for i := 2; i < len(dp); i++ {
		dp[i] = [][]int{}
		c := map[string]bool{}
		for _, num := range candidates {
			if i < num {
				break
			} else if i == num {
				dp[i] = append(dp[i], []int{num})
			} else {
				remain := i - num
				for _, nums := range dp[remain] {
					newCombination := append([]int{num}, nums...)
					sort.Ints(newCombination)
					key := makeKey(newCombination)
					if c[key] {
						continue
					} else {
						c[key] = true
						dp[i] = append(dp[i], newCombination)
					}
				}
			}
		}
	}
	return dp[target]
}
