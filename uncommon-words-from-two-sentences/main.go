package uncommonwordsfromtwosentences

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	countWords := map[string]int{}
	splitedS1 := strings.Split(s1, " ")
	splitedS2 := strings.Split(s2, " ")

	for _, s := range splitedS1 {
		countWords[s]++
	}
	for _, s := range splitedS2 {
		countWords[s]++
	}

	res := []string{}
	for s, v := range countWords {
		if v == 1 {
			res = append(res, s)
		}
	}
	return res
}
