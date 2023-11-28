package main

import "fmt"

// this is simply just calculate fibonacci at n position
func climbStairs(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		tmp := a
		a = b
		b += tmp
	}
	return b
}

func main() {
	fmt.Println(climbStairs(10))
}
