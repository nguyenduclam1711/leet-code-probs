package main

import "fmt"

func maxSubArray(nums []int) int {
	prefix := make([]int, len(nums))
	max := 0

	// calculate max prefix for each subarray
	for i, num := range nums {
		if i == 0 {
			prefix[i] = num
			max = prefix[i]
			continue
		}
		if num+prefix[i-1] > num {
			prefix[i] = num + prefix[i-1]
		} else {
			prefix[i] = num
		}
		if prefix[i] > max {
			max = prefix[i]
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
