package plusone

func PlusOne(digits []int) []int {
	needToPlus := true
	for i := len(digits) - 1; i >= 0; i-- {
		if needToPlus {
			digits[i]++
		}
		if digits[i] == 10 {
			digits[i] = 0
			needToPlus = true
		} else {
			needToPlus = false
		}
	}
	if digits[0] == 0 {
		return append([]int{1}, digits...)
	}
	return digits
}
