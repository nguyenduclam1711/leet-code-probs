package validpalindrome

import (
	"regexp"
	"strings"
)

var validAlphanumeric = regexp.MustCompile("[A-Za-z0-9]")

func checkIsAlphanumeric(s string) bool {
	return validAlphanumeric.MatchString(s)
}

func IsPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		str1, str2 := string(s[l]), string(s[r])

		if !checkIsAlphanumeric(str1) {
			l++
			continue
		}
		if !checkIsAlphanumeric(str2) {
			r--
			continue
		}
		if strings.ToLower(str1) == strings.ToLower(str2) {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}
