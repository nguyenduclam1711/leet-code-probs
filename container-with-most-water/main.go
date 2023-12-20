package containerwithmostwater

func findGreater(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxArea(height []int) int {
	area := 0
	l, r := 0, len(height)-1

	for l < r {
		width := r - l
		if height[l] < height[r] {
			area = findGreater(area, width*height[l])
			l++
		} else {
			area = findGreater(area, width*height[r])
			r--
		}
	}
	return area
}
