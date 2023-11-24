package main

import "fmt"

func findMinIndex(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mdl := (l + r) / 2

		if nums[mdl] > nums[r] {
			l = mdl + 1
		} else {
			r = mdl
		}
	}

	return l
}

func search(nums []int, target int) int {
	minIndex := findMinIndex(nums)
	l := 0
	r := len(nums) - 1

	if target == nums[minIndex] {
		return minIndex
	}
	if target > nums[minIndex] && target <= nums[r] {
		l = minIndex
	} else {
		r = minIndex - 1
	}
	for l < r {
		mdl := (l + r) / 2
		if target > nums[mdl] {
			l = mdl + 1
		} else {
			r = mdl
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func main() {
	fmt.Println(search([]int{1, 3}, 3))
}
