package issubsequence

func isSubsequence(s string, t string) bool {
	sPointer := 0

	for i := range t {
		if sPointer >= len(s) {
			return true
		}
		if t[i] == s[sPointer] {
			sPointer++
		}
	}
	return sPointer == len(s)
}
