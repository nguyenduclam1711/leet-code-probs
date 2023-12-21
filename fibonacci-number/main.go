package fibonaccinumber

func Fib(n int) int {
	if n < 2 {
		return n
	}
	s := []int{0, 1}

	for i := 2; i <= n; i++ {
		s = append(s, s[i-1]+s[i-2])
	}
	return s[len(s)-1]
}
