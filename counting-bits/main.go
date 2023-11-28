package main

import "fmt"

func countBits(n int) []int {
	len := n + 1
	res := make([]int, len)
	offset := 1

	for i := 1; i < len; i++ {
		if offset*2 == i {
			offset = i
		}
		res[i] = 1 + res[i-offset]
	}
	return res
}

func main() {
	fmt.Println(countBits(31))
}
