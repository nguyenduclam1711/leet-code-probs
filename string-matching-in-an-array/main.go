package stringmatchinginanarray

import "strings"

func stringMatching(words []string) []string {
	res := []string{}
	for _, word1 := range words {
		for _, word2 := range words {
			if word1 == word2 {
				continue
			}
			if strings.Contains(word2, word1) {
				res = append(res, word1)
				break
			}
		}
	}
	return res
}
