package wordpattern

import "strings"

func wordPattern(pattern string, s string) bool {
	splitedS := strings.Split(s, " ")

	if len(pattern) != len(splitedS) {
		return false
	}
	m := map[byte]string{}
	alrMapped := map[string]bool{}

	for i := range pattern {
		if _, e := m[pattern[i]]; !e {
			if alrMapped[splitedS[i]] {
				return false
			}
			m[pattern[i]] = splitedS[i]
			alrMapped[splitedS[i]] = true
		} else if m[pattern[i]] != splitedS[i] {
			return false
		}
	}
	return true
}
