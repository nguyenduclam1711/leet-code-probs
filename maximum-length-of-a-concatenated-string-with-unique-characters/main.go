package maximumlengthofaconcatenatedstringwithuniquecharacters

func findMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func recurse(currLen int, arr []string, visited [26]int) int {
	if len(arr) == 0 {
		return currLen
	}
	maxLen := currLen

	for i, s := range arr {
		isValid := true

		for _, c := range s {
			idx := c - 'a'

			if visited[idx] != 0 {
				isValid = false
			}
			visited[idx]++
		}

		if isValid {
			nextLen := recurse(currLen+len(s), arr[i+1:], visited)
			maxLen = findMax(maxLen, nextLen)
		}

		for _, c := range s {
			idx := c - 'a'
			visited[idx]--
		}
	}
	return maxLen
}

func maxLength(arr []string) int {
	visited := [26]int{}
	return recurse(0, arr, visited)
}
