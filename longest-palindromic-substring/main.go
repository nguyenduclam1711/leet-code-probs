package longestpalindromicsubstring

type PalindromePos struct {
	start int
	end   int
}

func LongestPalindrome(s string) string {
	queue := []PalindromePos{}
	maxLen := 0
	res := ""

	// init queue
	for i := range s {
		queue = append(queue, PalindromePos{i, i})
		if i < len(s)-1 {
			j := i + 1
			if s[i] == s[j] {
				queue = append(queue, PalindromePos{i, j})
			}
		}
	}
	for len(queue) > 0 {
		pos := queue[0]
		palindrome := s[pos.start : pos.end+1]

		if len(palindrome) > maxLen {
			res = palindrome
			maxLen = len(palindrome)
		}
		i, j := pos.start-1, pos.end+1
		if i >= 0 && j < len(s) && s[i] == s[j] {
			queue = append(queue, PalindromePos{i, j})
		}
		queue = queue[1:]
	}
	return res
}
