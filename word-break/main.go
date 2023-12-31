package wordbreak

func checkStrings(s1 string, wordFromDict string) (int, bool) {
	i := 0
	j := 0
	for i < len(s1) && j < len(wordFromDict) {
		if s1[i] == wordFromDict[j] {
			i++
			j++
		} else {
			return i, false
		}
	}
	if j == len(wordFromDict) {
		return i, true
	}
	return i, false
}

func recurse(s string, wordDict []string, caches map[string]bool) bool {
	res, exists := caches[s]
	if exists {
		return res
	}
	if len(s) == 0 {
		return true
	}
	validPos := []int{}
	for _, word := range wordDict {
		pos, isValid := checkStrings(s, word)
		if isValid {
			validPos = append(validPos, pos)
		}
	}
	result := false
	for _, p := range validPos {
		r := recurse(s[p:], wordDict, caches)
		if r {
			result = r
			break
		}
	}
	caches[s] = result
	return result
}

func WordBreak(s string, wordDict []string) bool {
	caches := map[string]bool{}
	return recurse(s, wordDict, caches)
}
