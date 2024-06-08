package isomorphicstrings

func isIsomorphic(s string, t string) bool {
	m := map[byte]byte{}
	alrMapped := map[byte]bool{}

	for i := range s {
		if _, e := m[s[i]]; !e {
			if alrMapped[t[i]] {
				return false
			}
			m[s[i]] = t[i]
			alrMapped[t[i]] = true
		} else if m[s[i]] != t[i] {
			return false
		}
	}
	return true
}
