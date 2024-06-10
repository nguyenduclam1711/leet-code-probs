package longestconsecutivesequence

func recurse(curr int, appears map[int]bool, cached map[int]int, currLen int) int {
	if cached[curr] > 0 {
		return cached[curr]
	}
	res := currLen
	nextN := curr + 1
	if appears[nextN] {
		res += recurse(nextN, appears, cached, 1)
	}
	cached[curr] = res
	return res
}

func longestConsecutive(nums []int) int {
	appears := map[int]bool{}
	for _, n := range nums {
		appears[n] = true
	}

	// store the length of consecutive by each num
	cached := map[int]int{}
	res := 0
	for _, n := range nums {
		currRes := recurse(n, appears, cached, 1)
		if currRes > res {
			res = currRes
		}
	}
	return res
}
