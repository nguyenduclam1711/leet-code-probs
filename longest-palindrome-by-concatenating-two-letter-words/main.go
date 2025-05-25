package longestpalindromebyconcatenatingtwoletterwords

func longestPalindrome(words []string) int {
	res := 0
	wordCount := map[string]int{}

	for _, w := range words {
		wordCount[w]++
	}

	leftPalindromeWords := []string{}

	for w := range wordCount {
		if wordCount[w] == 0 {
			continue
		}
		if w[0] == w[1] { // this word is already palindrome
			if wordCount[w]%2 == 0 {
				res += wordCount[w] * 2
				wordCount[w] = 0
			} else if wordCount[w] > 2 {
				res += (wordCount[w] / 2) * 4
				wordCount[w] = 1
				leftPalindromeWords = append(leftPalindromeWords, w)
			} else {
				leftPalindromeWords = append(leftPalindromeWords, w)
			}
			continue
		}
		otherPalindrome := string(w[1]) + string(w[0])
		if wordCount[otherPalindrome] > 0 {
			added := min(wordCount[w], wordCount[otherPalindrome])
			res += added * 4
			wordCount[w] -= added
			wordCount[otherPalindrome] -= added
		}
	}
	if len(leftPalindromeWords) > 0 {
		res += 2
	}
	return res
}
