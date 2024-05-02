package substringwithconcatenationofallwords

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
