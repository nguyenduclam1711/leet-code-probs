package main

import "fmt"

func removeElement(nums []int, val int) int {
	count := 0

	for _, num := range nums {
		if num != val {
			nums[count] = num
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
