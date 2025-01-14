package minimumlengthofstringafteroperations

func minimumLength(s string) int {
	numberOfChars := [26]int{}
	for _, r := range s {
		numberOfChars[r-'a']++
	}
	res := 0
	for _, num := range numberOfChars {
		if num == 0 {
			continue
		}
		if num%2 == 0 {
			res += 2
		} else {
			res += 1
		}
	}
	return res
}
