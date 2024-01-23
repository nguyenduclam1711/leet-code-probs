package greatestcommondivisorofstrings

func findDivisors(str string) []string {
	res := []string{str}
	maxLoop := len(str) / 2

	for i := 1; i <= maxLoop; i++ {
		if len(str)%i != 0 {
			continue
		}
		isValid := true
		for j := i; j < len(str); j += i {
			if str[:i] != str[j:j+i] {
				isValid = false
				break
			}
		}
		if isValid {
			res = append(res, str[:i])
		}
	}
	return res
}

func gcdOfStrings(str1 string, str2 string) string {
	divisors1 := findDivisors(str1)
	res := ""

	for _, s := range divisors1 {
		if len(str2)%len(s) != 0 {
			continue
		}
		isValid := true
		for j := 0; j < len(str2); j += len(s) {
			if s != str2[j:j+len(s)] {
				isValid = false
				break
			}
		}
		if isValid && len(s) > len(res) {
			res = s
		}
	}
	return res
}
