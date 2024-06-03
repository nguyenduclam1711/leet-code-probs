package minimumwindowsubstring

func checkValid(mapCountT map[byte]int, currCount map[byte]int) bool {
	res := true
	for k, v := range mapCountT {
		if currCount[k] < v {
			res = false
			break
		}
	}
	return res
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	mapCountT := map[byte]int{}
	for i := range t {
		mapCountT[t[i]]++
	}

	res := s + " "
	l, r := 0, 0
	currCount := map[byte]int{}
	for l < len(s) && r <= len(s) {
		isValid := checkValid(mapCountT, currCount)
		if !isValid {
			if r < len(s) {
				currCount[s[r]]++
			}
			r++
		} else {
			newRes := s[l:r]
			if len(newRes) < len(res) {
				res = newRes
			}
			currCount[s[l]]--
			l++
		}
	}
	if len(res) > len(s) {
		return ""
	}
	return res
}
