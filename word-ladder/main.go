package wordladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	uniqueWords := map[string]bool{}
	for _, w := range wordList {
		uniqueWords[w] = true
	}
	if !uniqueWords[endWord] {
		return 0
	}
	wLen := len(beginWord)
	q := []string{beginWord}
	checkedWords := map[string]bool{}
	checkedWords[beginWord] = true
	res := 0
	isValid := false
	for len(q) > 0 {
		res++
		newQ := []string{}
		for _, w := range q {
			if w == endWord {
				isValid = true
				break
			}
			for i := 0; i < wLen; i++ {
				for j := 0; j < 26; j++ {
					var nextW string
					currNewChar := string(rune('a' + j))
					if i == 0 {
						nextW = currNewChar + w[1:]
					} else if i == wLen-1 {
						nextW = w[:i] + currNewChar
					} else {
						nextW = w[:i] + currNewChar + w[i+1:]
					}
					if uniqueWords[nextW] && !checkedWords[nextW] {
						newQ = append(newQ, nextW)
						checkedWords[nextW] = true
					}
				}
			}
		}
		if isValid {
			break
		}
		q = newQ
	}
	if isValid {
		return res
	}
	return 0
}
