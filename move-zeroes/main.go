package main

import "fmt"

func moveZeroes(nums []int) {
	p := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
}

func main() {
	nums := []int{1, 1, 0, 3, 12, 0, 2, 3, 0, 8}
	moveZeroes(nums)
	fmt.Println(nums)
}
