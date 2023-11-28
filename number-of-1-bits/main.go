package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	w := 0

	for num != 0 {
		if num&1 == 1 {
			w++
		}
		num = num >> 1
	}
	return w
}

func main() {
	var num uint32 = 4294967293
	fmt.Println(hammingWeight(num))
}
