package findpivotindex

func pivotIndex(nums []int) int {
	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))

	for i := range nums {
		if i == 0 {
			leftSum[i] = 0
		} else {
			leftSum[i] = leftSum[i-1] + nums[i-1]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			rightSum[i] = 0
		} else {
			rightSum[i] = rightSum[i+1] + nums[i+1]
		}
	}
	for i := range nums {
		if leftSum[i] == rightSum[i] {
			return i
		}
	}
	return -1
}
