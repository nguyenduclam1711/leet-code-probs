package kokoeatingbananas

func minEatingSpeed(piles []int, h int) int {
	var maxNum int
	for _, p := range piles {
		if p > maxNum {
			maxNum = p
		}
	}
	if len(piles) == h {
		return maxNum
	}
	l, r := 1, maxNum
	for l < r {
		m := (l + r) / 2
		totalHours := 0
		for _, p := range piles {
			if m >= p {
				totalHours++
			} else {
				totalHours += p / m
				if p%m > 0 {
					totalHours++
				}
			}
		}
		if totalHours <= h {
			r = m
		} else {
			l = m + 1
		}
	}
	return r
}
