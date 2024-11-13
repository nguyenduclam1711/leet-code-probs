package largestcombinationwithbitwiseandgreaterthanzero

func find(n int, setBits []int) {
	i := 31

	for n > 0 {
		a := n & 1
		setBits[i] += a
		n >>= 1
		i--
	}
}

func largestCombination(candidates []int) int {
	setBits := make([]int, 32)
	for _, c := range candidates {
		find(c, setBits)
	}
	res := 0
	for _, newRes := range setBits {
		if newRes > res {
			res = newRes
		}
	}
	return res
}
