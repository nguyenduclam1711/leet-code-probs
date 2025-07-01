package countingbits

func CountBits(n int) []int {
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
