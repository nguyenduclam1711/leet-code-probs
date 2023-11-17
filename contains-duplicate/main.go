package main

import "fmt"

func containsDuplicate(nums []int) bool {
	var mapping = map[int]bool{}

	for _, num := range nums {
		if _, exists := mapping[num]; exists == true {
			return true
		} else {
			mapping[num] = true
		}
	}
	return false
}

func main() {
	fmt.Println("Input: nums = [1,2,3,1]")
	fmt.Println("Output:", containsDuplicate([]int{1, 2, 3, 1}))
}
