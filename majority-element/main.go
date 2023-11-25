package main

import "fmt"

func majorityElement(nums []int) int {
	appearTimes := len(nums) / 2
	m := map[int]int{}

	for _, num := range nums {
		m[num]++
		if m[num] > appearTimes {
			return num
		}
	}
	return -1
}

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
