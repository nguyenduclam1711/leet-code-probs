package generateparentheses

func generateParenthesis(n int) []string {
	openParenthesis := []rune{}
	closeParenthesis := []rune{}
	for i := 0; i < n; i++ {
		openParenthesis = append(openParenthesis, '(')
		closeParenthesis = append(closeParenthesis, ')')
	}
	res := []string{}
	var recurse func(s string)
	recurse = func(s string) {
		if len(openParenthesis) == 0 && len(closeParenthesis) == 0 {
			res = append(res, s)
			return
		}
		if len(openParenthesis) < len(closeParenthesis) {
			if len(openParenthesis) > 0 {
				// add (
				popOpen := openParenthesis[len(openParenthesis)-1]
				openParenthesis = openParenthesis[:len(openParenthesis)-1]
				recurse(s + string(popOpen))
				openParenthesis = append(openParenthesis, '(')
			}

			if len(closeParenthesis) > 0 {
				// add )
				popClose := closeParenthesis[len(closeParenthesis)-1]
				closeParenthesis = closeParenthesis[:len(closeParenthesis)-1]
				recurse(s + string(popClose))
				closeParenthesis = append(closeParenthesis, ')')
			}
		} else if len(openParenthesis) == len(closeParenthesis) {
			if len(openParenthesis) > 0 {
				// add (
				popOpen := openParenthesis[len(openParenthesis)-1]
				openParenthesis = openParenthesis[:len(openParenthesis)-1]
				recurse(s + string(popOpen))
				openParenthesis = append(openParenthesis, '(')
			}
		}
	}
	recurse("")
	return res
}
