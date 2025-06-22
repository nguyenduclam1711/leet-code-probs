package keyskeyboard

var primeNumbers = []int{2, 3, 5, 7}

func checkIsPrimeNumber(n int) (bool, int) {
	result := true
	largestDivisible := 0
	for i := range primeNumbers {
		primeNumber := primeNumbers[i]
		if n == primeNumber {
			return result, largestDivisible
		}
		if n%primeNumber == 0 {
			result = false
			largestDivisible = n / primeNumber
			break
		}
	}
	if result {
		primeNumbers = append(primeNumbers, n)
	}
	return result, largestDivisible
}

func minSteps(n int) int {
	dp := make([]int, n+1)
	if n > 1 {
		dp[2] = 2
	}
	for i := 3; i < len(dp); i++ {
		isPrimeNumber, largestDivisible := checkIsPrimeNumber(i)
		if isPrimeNumber {
			dp[i] = i
		} else {
			dp[i] = dp[largestDivisible] + i/largestDivisible
		}
	}
	return dp[n]
}
