package main

import "fmt"

func singleNumber(nums []int) int {
	m := map[int]bool{}

	for _, num := range nums {
		if m[num] {
			delete(m, num)
		} else {
			m[num] = true
		}
	}
	for key := range m {
		return key
	}
	return 0
}

func main() {
	fmt.Println([]int{4, 1, 2, 1, 2})
}
