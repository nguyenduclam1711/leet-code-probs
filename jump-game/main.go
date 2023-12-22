package jumpgame

func calculateDistanceNeedToReach(canJumpPos []int, currIdx int) int {
	posNeedToReach := canJumpPos[len(canJumpPos)-1]
	return posNeedToReach - currIdx
}

func CanJump(nums []int) bool {
	canJumpPos := []int{len(nums) - 1}

	for i := len(nums) - 2; i >= 0; i-- {
		dist := calculateDistanceNeedToReach(canJumpPos, i)
		maxJumpLength := nums[i]

		if maxJumpLength >= dist {
			canJumpPos = append(canJumpPos, i)
		}
	}
	return canJumpPos[len(canJumpPos)-1] == 0
}
