package validparentheses

var mapping = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func IsValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if mapping[c] == stack[len(stack)-1] {
			stack = stack[0 : len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
