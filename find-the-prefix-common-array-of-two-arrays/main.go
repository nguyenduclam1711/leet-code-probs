package findtheprefixcommonarrayoftwoarrays

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	m := make([]int, n+1)
	res := []int{}
	for i := 0; i < n; i++ {
		a, b := A[i], B[i]
		m[a]++
		m[b]++
		curr := 0
		if len(res) > 0 {
			curr = res[i-1]
		}
		if a == b {
			if m[a] == 2 {
				curr++
			}
		} else {
			if m[a] == 2 {
				curr++
			}
			if m[b] == 2 {
				curr++
			}
		}
		res = append(res, curr)
	}
	return res
}
