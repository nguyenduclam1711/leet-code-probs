package maximumnumberofvowelsinasubstringofgivenlength

func maxVowels(s string, k int) int {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	l, r := 0, 0
	currNumberOfVowels := 0
	res := 0

	for l < len(s) && r < len(s) {
		windowLen := r - l + 1

		if vowels[s[r]] {
			currNumberOfVowels++
		}
		if windowLen > k {
			if vowels[s[l]] {
				currNumberOfVowels--
			}
			l++
		}
		if currNumberOfVowels > res {
			res = currNumberOfVowels
		}
		r++
	}
	return res
}
