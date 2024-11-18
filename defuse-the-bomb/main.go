package defusethebomb

func decrypt(code []int, k int) []int {
	if k == 0 {
		return make([]int, len(code))
	}
	var l, r int

	if k > 0 {
		l, r = 1, k
	} else {
		l, r = len(code)+k, len(code)-1
	}
	currSum := code[l]
	for i := l + 1; i <= r; i++ {
		currSum += code[i]
	}
	res := make([]int, len(code))
	res[0] = currSum

	for i := 1; i < len(res); i++ {
		currSum -= code[l]
		l = (l + 1) % len(res)
		r = (r + 1) % len(res)
		currSum += code[r]
		res[i] = currSum
	}
	return res
}
