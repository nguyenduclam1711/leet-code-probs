package minimumwindowsubstring

func checkValid(mapCountT map[byte]int, currCount map[byte]int, leftChars map[byte]bool) bool {
	res := true
	for char := range leftChars {
		if mapCountT[char] > currCount[char] {
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
		isValid := checkValid(mapCountT, currCount, leftChars)
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
			if _, e := mapCountT[s[l]]; e {
				leftChars[s[l]] = true
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
