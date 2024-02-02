package removingstarsfromastring

func removeStars(s string) string {
	bytes := []byte{}

	for i := range s {
		if s[i] != '*' {
			bytes = append(bytes, s[i])
		} else {
			bytes = bytes[:len(bytes)-1]
		}
	}
	return string(bytes)
}
