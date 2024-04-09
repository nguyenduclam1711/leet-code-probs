package thekthfactorofn

func kthFactor(n int, k int) int {
	factors := []int{}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
		if len(factors) == k {
			return factors[len(factors)-1]
		}
	}
	return -1
}
