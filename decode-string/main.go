package decodestring

import (
	"strconv"
)

func decodeString(s string) string {
	nums := []string{}
	chars := []string{}
	res := ""

	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			if len(nums) == 0 || (i > 0 && (s[i-1] < '0' || s[i-1] > '9')) {
				nums = append(nums, string(s[i]))
			} else {
				nums[len(nums)-1] += string(s[i])
			}
		} else if s[i] == '[' {
			if len(nums) > len(chars) {
				chars = append(chars, "")
			}
		} else if s[i] >= 'a' && s[i] <= 'z' {
			if len(chars) == 0 || len(chars) < len(nums) {
				chars = append(chars, string(s[i]))
			} else {
				chars[len(chars)-1] += string(s[i])
			}
			if len(chars) > len(nums) {
				nums = append(nums, string('1'))
			}
		} else if s[i] == ']' {
			numOfChars, _ := strconv.Atoi(nums[len(nums)-1])
			if len(nums) == 1 {
				for i := 0; i < numOfChars; i++ {
					res += chars[len(chars)-1]
				}
			} else if len(nums) > 1 {
				for i := 0; i < numOfChars; i++ {
					chars[len(chars)-2] += chars[len(chars)-1]
				}
			}
			chars = chars[:len(chars)-1]
			nums = nums[:len(nums)-1]
		}
	}
	if len(chars) > 0 {
		res += chars[0]
	}
	return res
}
