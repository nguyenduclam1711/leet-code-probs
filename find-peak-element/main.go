package findpeakelement

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		adjL, adjR := m-1, m+1
		isValidL := adjL < 0 || nums[adjL] < nums[m]
		isValidR := adjR > len(nums)-1 || nums[adjR] < nums[m]
		if isValidL && isValidR {
			return m
		} else if !isValidL {
			r = m
		} else {
			l = m + 1
		}
	}
	return -1
}
