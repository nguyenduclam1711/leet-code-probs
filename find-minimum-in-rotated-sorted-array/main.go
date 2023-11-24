package main

import (
	"fmt"
)

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mdl := (l + r) / 2

		if nums[mdl] > nums[r] {
			l = mdl + 1
		} else {
			r = mdl
		}
	}

	return nums[l]
}

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 8, 9, 10, 0, 1, 2, 3}))
}
