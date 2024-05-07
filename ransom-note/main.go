package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}
	hashMap := map[rune]int{}

	for _, r := range magazine {
		hashMap[r]++
	}
	for _, r := range ransomNote {
		if hashMap[r] <= 0 {
			return false
		}
		hashMap[r]--
	}
	return true
}
