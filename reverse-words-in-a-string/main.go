package reversewordsinastring

func reverseWords(s string) string {
	words := []string{}

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		runes := []byte{s[i]}

		j := i + 1
		for ; j < len(s); j++ {
			if s[j] != ' ' {
				runes = append(runes, s[j])
			} else {
				break
			}
		}
		words = append(words, string(runes))
		i = j
	}

	res := ""
	for i := len(words) - 1; i >= 0; i-- {
		res += words[i]
		if i > 0 {
			res += " "
		}
	}
	return res
}
