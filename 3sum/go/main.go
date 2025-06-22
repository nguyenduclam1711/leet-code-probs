package threesum

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	var prevA int

	for i, num := range nums {
		if num == prevA && i != 0 {
			continue
		}
		l, r := i+1, len(nums)-1
		// init prevB like this because this initial will never includes in nums (nums is sorted so I can do this)
		prevB := nums[0] - 1

		for l < r {
			if nums[l] == prevB {
				l++
				continue
			}
			tmp := nums[l] + nums[r]
			if tmp == -num {
				result = append(result, []int{num, nums[l], nums[r]})
				prevB = nums[l]
				l++
				r--
			} else if tmp < -num {
				l++
			} else {
				r--
			}
		}
		prevA = num
	}
	return result
}
