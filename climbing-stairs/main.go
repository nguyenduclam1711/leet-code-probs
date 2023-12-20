package climbingstairs

// this is simply just calculate fibonacci at n position
func ClimbStairs(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		tmp := a
		a = b
		b += tmp
	}
	return b
}
