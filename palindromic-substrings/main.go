package palindromicsubstrings

type PalindromePos struct {
	start int
	end   int
}

func addToQueue(q *[]PalindromePos, start int, end int, count *int) {
	*q = append(*q, PalindromePos{start, end})
	*count++
}

func CountSubstrings(s string) int {
	queue := []PalindromePos{}
	count := 0

	// init queue
	for i := range s {
		addToQueue(&queue, i, i, &count)
		if i < len(s)-1 {
			j := i + 1
			if s[i] == s[j] {
				addToQueue(&queue, i, j, &count)
			}
		}
	}
	for len(queue) > 0 {
		pos := queue[0]
		i, j := pos.start-1, pos.end+1
		if i >= 0 && j < len(s) && s[i] == s[j] {
			addToQueue(&queue, i, j, &count)
		}
		queue = queue[1:]
	}
	return count
}
