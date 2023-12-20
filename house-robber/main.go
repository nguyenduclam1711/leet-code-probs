package houserobber

func findMax(nums []int) int {
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func Rob(nums []int) int {
	max := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 || i == len(nums)-2 {
			max[i] = nums[i]
		} else {
			max[i] = nums[i] + findMax(max[i+2:])
		}
	}
	if len(max) == 1 {
		return max[0]
	}
	if max[0] > max[1] {
		return max[0]
	}
	return max[1]
}
