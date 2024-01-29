package maxnumberofksumpairs

import "sort"

// this is using 2 pointers to solve
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

// using hashmap
func maxOperations2(nums []int, k int) int {
	m := map[int]int{}
	res := 0

	for _, n := range nums {
		dif := k - n

		if m[n] > 0 {
			m[n]--
			res++
		} else {
			m[dif]++
		}
	}
	return res
}
