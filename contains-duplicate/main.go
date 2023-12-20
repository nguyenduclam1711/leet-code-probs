package containsduplicate

func ContainsDuplicate(nums []int) bool {
	mapping := map[int]bool{}

	for _, num := range nums {
		if _, exists := mapping[num]; exists {
			return true
		} else {
			mapping[num] = true
		}
	}
	return false
}
