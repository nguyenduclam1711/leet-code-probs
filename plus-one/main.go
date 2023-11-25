package main

import "fmt"

func plusOne(digits []int) []int {
	needToPlus := true
	for i := len(digits) - 1; i >= 0; i-- {
		if needToPlus {
			digits[i]++
		}
		if digits[i] == 10 {
			digits[i] = 0
			needToPlus = true
		} else {
			needToPlus = false
		}
	}
	if digits[0] == 0 {
		return append([]int{1}, digits...)
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{9}))
}
