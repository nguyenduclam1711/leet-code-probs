package main

import "fmt"

func removeDuplicates(nums []int) int {
	count := 1
	prevNum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != prevNum {
			nums[count] = nums[i]
			count++
			prevNum = nums[i]
		}
	}
	return count
}

func main() {
	fmt.Println(removeDuplicates([]int{2, 3, 2, 3}))
}
