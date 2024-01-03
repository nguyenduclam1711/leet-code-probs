package longestsubstringwithoutrepeatingcharacters

func LengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	l, r := 0, 1
	maxLen := 1
	hashMap := map[string]bool{
		string(s[l]): true,
	}

	for l < len(s) && r < len(s) {
		if l == r {
			r = l + 1
			hashMap[string(s[l])] = true
			continue
		}
		if !hashMap[string(s[r])] {
			hashMap[string(s[r])] = true
			currLen := r - l + 1

			if currLen > maxLen {
				maxLen = currLen
			}
			r++
		} else {
			hashMap[string(s[l])] = false
			l++
		}
	}
	return maxLen
}
