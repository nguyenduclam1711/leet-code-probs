package dailytemperatures

func dailyTemperatures(temperatures []int) []int {
	stack := [][2]int{
		{len(temperatures) - 1, temperatures[len(temperatures)-1]},
	}
	res := make([]int, len(temperatures))

	for i := len(temperatures) - 2; i >= 0; i-- {
		curTemp := temperatures[i]
		for len(stack) > 0 {
			if curTemp >= stack[len(stack)-1][1] {
				stack = stack[:len(stack)-1]
			} else {
				res[i] = stack[len(stack)-1][0] - i
				stack = append(stack, [2]int{i, curTemp})
				break
			}
		}
		if len(stack) == 0 {
			stack = append(stack, [2]int{i, curTemp})
		}
	}
	return res
}
