package main

import "fmt"

func hammingWeight(n int) int {
	w := 0
	for n != 0 {
		if n&1 == 1 {
			w++
		}
		n = n >> 1
	}
	return w
}

func countBits(n int) []int {
	res := make([]int, n+1)

	for i := range res {
		res[i] = hammingWeight(i)
	}
	return res
}

func main() {
	fmt.Println(countBits(31))
}
