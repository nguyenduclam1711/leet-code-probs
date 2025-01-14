package uglynumberii

func nthUglyNumber(n int) int {
	ugly := make([]int, n)
	ugly[0] = 1
	index1, index2, index3 := 0, 0, 0
	factor1, factor2, factor3 := 2, 3, 5

	for i := 1; i < n; i++ {
		nextUgly := min(factor1, factor2, factor3)
		ugly[i] = nextUgly
		if factor1 == nextUgly {
			index1++
			factor1 = 2 * ugly[index1]
		}
		if factor2 == nextUgly {
			index2++
			factor2 = 3 * ugly[index2]
		}
		if factor3 == nextUgly {
			index3++
			factor3 = 5 * ugly[index3]
		}
	}
	return ugly[n-1]
}
