package threesumclosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minDiff := math.MaxInt
	result := 0

	for i, n := range nums {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := n + nums[l] + nums[r]
			if sum == target {
				return sum
			}
			diff := int(math.Abs(float64(target - sum)))
			if diff < minDiff {
				minDiff = diff
				result = sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return result
}
