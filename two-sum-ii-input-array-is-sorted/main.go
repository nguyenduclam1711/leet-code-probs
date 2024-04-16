package twosumiiinputarrayissorted

func jump(nums []int) int {
	i := 0
	res := 0
	for i < len(nums)-1 {
		res++
		n := nums[i]
		if i+n >= len(nums)-1 {
			return res
		}
		maxJumpPos, nextPos := -1, i
		for j := i + 1; j <= n+i; j++ {
			curJumpPos := j + nums[j]
			if curJumpPos > maxJumpPos {
				maxJumpPos = curJumpPos
				nextPos = j
			}
		}
		i = nextPos
	}
	return res
}
