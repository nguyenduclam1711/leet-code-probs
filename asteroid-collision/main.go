package asteroidcollision

import "math"

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, n := range asteroids {
		stack = append(stack, n)
		for len(stack) > 1 {
			r := len(stack) - 1
			l := r - 1
			if stack[l] > 0 && stack[r] < 0 {
				rAbs := math.Abs(float64(stack[r]))
				lAbs := math.Abs(float64(stack[l]))
				if rAbs == lAbs {
					stack = stack[:l]
				} else {
					if rAbs > lAbs {
						stack[l] = stack[r]
					}
					stack = stack[:r]
				}
			} else {
				break
			}
		}
	}
	return stack
}
