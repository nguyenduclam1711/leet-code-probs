package containsduplicate

import "sort"

func ContainsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := range len(nums) - 1 {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
