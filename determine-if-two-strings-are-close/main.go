package determineiftwostringsareclose

import "sort"

func getMapChars(word string) []int {
	mapChars := make([]int, 26)
	for i := range word {
		mapChars[word[i]-'a']++
	}
	return mapChars
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	mapChars1 := getMapChars(word1)
	mapChars2 := getMapChars(word2)

	for i := range mapChars1 {
		if mapChars1[i] > 0 && mapChars2[i] == 0 {
			return false
		}
	}

	sort.Ints(mapChars1)
	sort.Ints(mapChars2)

	for i := range mapChars1 {
		if mapChars2[i] != mapChars1[i] {
			return false
		}
	}
	return true
}
