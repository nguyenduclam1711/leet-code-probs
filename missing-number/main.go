package main

import "fmt"

func missingNumber(nums []int) int {
	totalInNumsShouldBe := len(nums) * (len(nums) + 1) / 2
	totalInNums := 0
	for _, num := range nums {
		totalInNums += num
	}
	return totalInNumsShouldBe - totalInNums
}

func main() {
	fmt.Println(missingNumber([]int{0, 1, 3}))
}
