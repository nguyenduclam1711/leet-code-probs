package uniquenumberofoccurrences

func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}
	for _, n := range arr {
		m[n]++
	}

	occur := map[int]int{}
	for k, v := range m {
		if _, e := occur[v]; !e {
			occur[v] = k
		} else {
			return false
		}
	}
	return true
}
