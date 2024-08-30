package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	res := false
	for _, row := range matrix {
		last := row[len(row)-1]

		if target > last {
			continue
		}
		l, r := 0, len(row)

		for l < r {
			m := (l + r) / 2
			if target == row[m] {
				res = true
				break
			} else if target > row[m] {
				l = m + 1
			} else {
				r = m
			}
		}
		break
	}
	return res
}
