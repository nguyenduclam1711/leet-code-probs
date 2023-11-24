package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)

	for i, num := range nums {
		if _, exist := mapping[target-num]; exist {
			return []int{mapping[target-num], i}
		}
		mapping[num] = i
	}
	return []int{}
}

func main() {
	fmt.Println("Input: nums = [2,7,11,15], target = 9")
	fmt.Println("Output:", twoSum([]int{2, 7, 11, 15}, 9))
}
