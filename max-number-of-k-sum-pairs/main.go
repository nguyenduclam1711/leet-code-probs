package maxnumberofksumpairs

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	l := 0
	r := len(nums) - 1
	res := 0

	for l < r {
		dif := k - nums[l]

		if dif == nums[r] {
			l++
			r--
			res++
		} else if dif > nums[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
