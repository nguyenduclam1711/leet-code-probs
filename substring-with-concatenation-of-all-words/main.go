package substringwithconcatenationofallwords

// using hash map and clone hash map for checking
func findSubstring(s string, words []string) []int {
	hashMap := map[string]int{}
	wordLen := len(words[0])
	windowLength := len(words) * wordLen
	res := []int{}

	for _, word := range words {
		hashMap[word]++
	}
	for i := 0; i <= len(s)-windowLength; i++ {
		// make new hashmap for tracking
		currHashMap := map[string]int{}
		for k, v := range hashMap {
			currHashMap[k] = v
		}

		j := i
		for ; j < i+windowLength; j += wordLen {
			currWord := s[j : j+wordLen]
			if currHashMap[currWord] <= 0 {
				break
			}
			currHashMap[currWord]--
		}
		if j == i+windowLength {
			res = append(res, i)
		}
	}
	return res
}

// still use hashmap but not clone hashmap from the beginning => reduce the space uses
func findSubstring2(s string, words []string) []int {
	hashMap := map[string]int{}
	for _, word := range words {
		hashMap[word]++
	}
	wordLen := len(words[0])
	windowLen := len(words) * wordLen
	res := []int{}

	for i := 0; i <= len(s)-windowLen; i++ {
		seenWords := map[string]int{}
		count := 0

		for j := i; j < i+windowLen; j += wordLen {
			currWord := s[j : j+wordLen]
			if _, e := hashMap[currWord]; e {
				seenWords[currWord]++
				count++

				if seenWords[currWord] > hashMap[currWord] {
					count = 0
					break
				}
			} else {
				count = 0
				break
			}
		}
		if count == len(words) {
			res = append(res, i)
		}
	}
	return res
}
