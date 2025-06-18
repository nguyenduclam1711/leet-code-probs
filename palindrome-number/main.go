package palindromenumber

import "fmt"

func isPalindrome(x int) bool {
	str := fmt.Sprint(x)
	l, r := 0, len(str)-1

	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}
