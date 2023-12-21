package uniquepaths

import "fmt"

func recurse(m int, n int, x int, y int, cache map[string]int) int {
	s := fmt.Sprint(x, ",", y)
	_, exists := cache[s]
	if exists {
		return cache[s]
	}
	if x == n || y == m {
		cache[s] = 1
		return 1
	}
	res := recurse(m, n, x+1, y, cache) + recurse(m, n, x, y+1, cache)
	cache[s] = res
	return res
}

func UniquePaths(m int, n int) int {
	cache := map[string]int{}
	return recurse(m, n, 1, 1, cache)
}
