package removeduplicatesfromsortedarrayii

func removeDuplicates(nums []int) int {
	pos := 1
	apps := 1
	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == curr {
			if apps < 2 {
				apps++
				if pos < i {
					nums[pos] = nums[i]
				}
				pos++
			}
		} else {
			curr = nums[i]
			apps = 1
			if pos < i {
				nums[pos] = nums[i]
			}
			pos++
		}
	}
	return pos
}
