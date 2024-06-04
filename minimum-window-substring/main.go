package minimumwindowsubstring

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	mapCountT := map[byte]int{}
	leftChars := map[byte]bool{}
	for i := range t {
		if _, e := mapCountT[t[i]]; !e {
			leftChars[t[i]] = true
		}
		mapCountT[t[i]]++
	}

	res := s + " "
	l, r := 0, 0
	currCount := map[byte]int{}
	for l < len(s) && r <= len(s) {
		isValid := len(leftChars) == 0
		if !isValid {
			if r < len(s) {
				currCount[s[r]]++
				if currCount[s[r]] >= mapCountT[s[r]] {
					delete(leftChars, s[r])
				}
			}
			r++
		} else {
			newRes := s[l:r]
			if len(newRes) < len(res) {
				res = newRes
			}
			currCount[s[l]]--
			if mapCountT[s[l]] > 0 && currCount[s[l]] < mapCountT[s[l]] {
				leftChars[s[l]] = true
			}
			l++
		}
	}
	if len(res) > len(s) {
		return ""
	}
	return res
}
