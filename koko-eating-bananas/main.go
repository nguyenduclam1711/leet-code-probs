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

	l, r := 1, maxNum+1
	res := 0
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
		if totalHours == h {
			res = m
			break
		} else if totalHours < h {
			res = m
			r = m
		} else {
			l = m + 1
		}
	}
	// still need to check if there are any value that lower than the res above that still make the total hours equal to h
	for res > 1 {
		nextRes := res - 1
		totalHours := 0
		for _, p := range piles {
			if nextRes >= p {
				totalHours++
			} else {
				totalHours += p / nextRes
				if p%nextRes > 0 {
					totalHours++
				}
			}
		}
		if totalHours == h {
			res = nextRes
		} else {
			break
		}
	}
	return res
}
