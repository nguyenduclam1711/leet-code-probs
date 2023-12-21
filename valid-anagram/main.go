package validanagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dict1 := map[byte]int{}
	dict2 := map[byte]int{}

	for i := 0; i < len(s); i++ {
		dict1[s[i]]++
		dict2[t[i]]++
	}
	if len(dict1) != len(dict2) {
		return false
	}
	for key, val := range dict1 {
		if dict2[key] != val {
			return false
		}
	}
	return true
}
