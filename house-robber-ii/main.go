package houserobberii

func findMax(nums []int) int {
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// from house robber 1 problem
func rob1(nums []int) int {
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

func Rob(nums []int) int {
	if len(nums) <= 3 {
		return findMax(nums)
	}
	max := 0

	for i, num := range nums {
		start := (i + len(nums) + 2) % len(nums)
		end := (i + len(nums) - 2) % len(nums)
		tmp := num

		if start == end {
			tmp += nums[start]
		} else if start < end {
			tmp += rob1(nums[start : end+1])
		} else {
			tmp += rob1(append(nums[start:], nums[:end+1]...))
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}
