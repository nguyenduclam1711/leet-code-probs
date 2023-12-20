package singlenumber

func SingleNumber(nums []int) int {
	m := map[int]bool{}

	for _, num := range nums {
		if m[num] {
			delete(m, num)
		} else {
			m[num] = true
		}
	}
	for key := range m {
		return key
	}
	return 0
}
