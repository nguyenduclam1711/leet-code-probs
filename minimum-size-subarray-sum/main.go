package minimumsizesubarraysum

func minSubArrayLen(target int, nums []int) int {
	windowLength := len(nums) + 1
	p1, p2 := 0, 0
	currSum := 0 - nums[p1]

	for p1 <= p2 && p1 < len(nums) && p2 < len(nums) {
		currSum += nums[p1] + nums[p2]
		if currSum < target {
			currSum -= nums[p1]
			p2++
		} else {
			currLength := p2 - p1 + 1
			if currLength < windowLength {
				windowLength = currLength
			}
			if currSum == target {
				currSum -= nums[p1]
				p1++
				if p1 > p2 || p1 >= len(nums) {
					break
				}
				currSum -= nums[p1]
				p2++
			} else {
				currSum -= nums[p1]
				p1++
				if p1 > p2 || p1 >= len(nums) {
					break
				}
				currSum -= nums[p1]
				currSum -= nums[p2]
			}
		}
	}
	if windowLength == len(nums)+1 {
		return 0
	}
	return windowLength
}
