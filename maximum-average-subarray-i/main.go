package maximumaveragesubarrayi

import "math"

func findMaxAverage(nums []int, k int) float64 {
	i, j := 0, 0
	totalMax := math.MinInt
	var currTotal int

	for j < len(nums) {
		length := j - i + 1
		currTotal += nums[j]

		if length > k {
			currTotal -= nums[i]
			i++
			length--
		}
		if length == k && currTotal > totalMax {
			totalMax = currTotal
		}
		j++
	}
	return float64(totalMax) / float64(k)
}
