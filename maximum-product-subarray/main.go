package main

import (
	"fmt"
)

func findMin(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func findMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProduct(nums []int) int {
	result := 0
	min := 0
	max := 0

	for i, num := range nums {
		if num == 0 {
			if num > result {
				result = num
			}
			min = 1
			max = 1
			continue
		}
		if i == 0 {
			min = num
			max = num
			result = num
			continue
		}
		x := min * num
		y := max * num

		min = findMin(findMin(x, y), num)
		max = findMax(findMax(x, y), num)

		if max > result {
			result = max
		}
	}
	return result
}

func main() {
	fmt.Println(maxProduct([]int{2, -5, -2, -4, 3}))
}
