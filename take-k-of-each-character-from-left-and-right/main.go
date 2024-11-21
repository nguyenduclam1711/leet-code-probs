package takekofeachcharacterfromleftandright

func takeCharacters(s string, k int) int {
	countChars := make([]int, 3)
	for _, c := range s {
		countChars[c-'a']++
	}
	if countChars[0] < k || countChars[1] < k || countChars[2] < k {
		return -1
	}
	res := len(s)
	var l, r, windowSize int
	currCountChars := make([]int, 3)
	for r < len(s) {
		currCountChars[s[r]-'a']++
		windowSize = r - l + 1
		countLeftA := countChars[0] - currCountChars[0]
		countLeftB := countChars[1] - currCountChars[1]
		countLeftC := countChars[2] - currCountChars[2]
		if countLeftA >= k && countLeftB >= k && countLeftC >= k {
			newRes := len(s) - windowSize
			if newRes < res {
				res = newRes
			}
		} else {
			currCountChars[s[l]-'a']--
			l++
		}
		r++
	}
	return res
}
