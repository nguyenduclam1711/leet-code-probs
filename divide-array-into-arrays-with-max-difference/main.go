package dividearrayintoarrayswithmaxdifference

import (
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	i := 0

	for i < len(nums) {
		if nums[i+2]-nums[i] <= k {
			res = append(res, nums[i:i+3])
			i += 3
		} else {
			return [][]int{}
		}
	}
	return res
}
