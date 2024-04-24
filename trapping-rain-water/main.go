package trappingrainwater

func trap(height []int) int {
	if len(height) == 1 {
		return 0
	}
	relativeTopsStack := []int{}
	res := 0

	for i, h := range height {
		if h == 0 {
			continue
		}
		if len(relativeTopsStack) == 0 {
			relativeTopsStack = append(relativeTopsStack, i)
			continue
		}
		lastTopPos := relativeTopsStack[len(relativeTopsStack)-1]
		lastH := height[lastTopPos]

		for len(relativeTopsStack) > 0 {
			if h <= lastH {
				relativeTopsStack = append(relativeTopsStack, i)
				break
			}
			relativeTopsStack = relativeTopsStack[:len(relativeTopsStack)-1]
			if len(relativeTopsStack) > 0 {
				lastTopPos = relativeTopsStack[len(relativeTopsStack)-1]
				lastH = height[lastTopPos]
			}
		}
		if len(relativeTopsStack) == 0 {
			relativeTopsStack = append(relativeTopsStack, i)
			// calculate amount of water
			minH := min(lastH, h)
			for j := lastTopPos + 1; j < i; j++ {
				res += minH - height[j]
			}
		}
	}
	// recalculate amount of water if there the length of the stack > 1
	if len(relativeTopsStack) > 1 {
		for i := 0; i < len(relativeTopsStack)-1; i++ {
			currPos, nextPos := relativeTopsStack[i], relativeTopsStack[i+1]

			if nextPos-currPos <= 1 {
				continue
			}
			currH, nextH := height[currPos], height[nextPos]
			minH := min(currH, nextH)
			for j := currPos + 1; j < nextPos; j++ {
				res += minH - height[j]
			}
		}
	}
	return res
}
