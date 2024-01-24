package reversevowelsofastring

func reverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	i := 0
	j := len(s) - 1
	runes := make([]byte, len(s))

	for i <= j {
		if !vowels[s[i]] {
			runes[i] = s[i]
			i++
			continue
		}
		if !vowels[s[j]] {
			runes[j] = s[j]
			j--
			continue
		}
		runes[i] = s[j]
		runes[j] = s[i]
		i++
		j--
	}
	return string(runes)
}
