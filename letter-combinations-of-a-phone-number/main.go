package lettercombinationsofaphonenumber

func dfs(i int, digits string, chars *[]byte, charsMap map[byte][]byte, res *[]string) {
	digit := digits[i]
	for _, char := range charsMap[digit] {
		*chars = append(*chars, char)
		if i == len(digits)-1 {
			*res = append(*res, string(*chars))
		} else {
			dfs(i+1, digits, chars, charsMap, res)
		}
		*chars = (*chars)[:len(*chars)-1]
	}
}

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	charsMap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	chars := []byte{}
	dfs(0, digits, &chars, charsMap, &res)
	return res
}
