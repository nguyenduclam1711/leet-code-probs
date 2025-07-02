package longestharmonioussubsequence

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// couting
func findLHS(nums []int) int {
	count := map[string]int{}
	uniqueNums := map[int]bool{}

	for _, n := range nums {
		uniqueNums[n] = true
		combination := fmt.Sprint(n, " ", n+1)
		count[combination]++
		combination2 := fmt.Sprint(n-1, " ", n)
		count[combination2]++
	}
	if len(uniqueNums) == 1 {
		return 0
	}
	res := 1
	for k, val := range count {
		if val > 1 {
			splitedKeys := strings.Split(k, " ")
			n1, _ := strconv.Atoi(splitedKeys[0])
			n2, _ := strconv.Atoi(splitedKeys[1])
			if uniqueNums[n1] && uniqueNums[n2] {
				res = max(res, val)
			}
		}
	}
	if res == 1 {
		return 0
	}
	return res
}

// sort + sliding window
func findLHS2(nums []int) int {
	sort.Ints(nums)
	res := 0
	j := 0

	for i := range nums {
		for nums[i]-nums[j] > 1 {
			j++
		}
		if nums[i]-nums[j] == 1 {
			res = max(res, i-j+1)
		}
	}
	return res
}
