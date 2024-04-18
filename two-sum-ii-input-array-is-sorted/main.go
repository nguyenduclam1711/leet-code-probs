package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	res := []int{}
	for i := 0; i < len(numbers)-1; i++ {
		num1 := numbers[i]
		left := target - num1
		done := false
		// binary search
		l, r := i+1, len(numbers)

		for l < r {
			m := (l + r) / 2

			if left == numbers[m] {
				res = append(res, i+1, m+1)
				done = true
				break
			} else if numbers[m] > left {
				r = m
			} else {
				l = m + 1
			}
		}
		if done {
			break
		}
	}
	return res
}
