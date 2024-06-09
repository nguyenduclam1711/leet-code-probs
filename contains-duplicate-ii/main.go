package containsduplicateii

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}

	for i, n := range nums {
		if pos, e := m[n]; !e {
			m[n] = i
		} else {
			diffAbs := math.Abs(float64(i - pos))
			if int(diffAbs) <= k {
				return true
			}
			m[n] = i
		}
	}
	return false
}
