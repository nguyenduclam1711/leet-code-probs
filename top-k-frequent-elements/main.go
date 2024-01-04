package topkfrequentelements

func TopKFrequent(nums []int, k int) []int {
	s := make([][]int, len(nums)+1)
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		s[v] = append(s[v], k)
	}

	result := []int{}
	i := k
	for j := len(s) - 1; j > 0; j-- {
		numbers := s[j]
		for _, num := range numbers {
			if i == 0 {
				break
			}
			result = append(result, num)
			i--
		}
	}
	return result
}
