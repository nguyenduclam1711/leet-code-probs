package zigzagconversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	arr := make([][]byte, numRows)
	for i := range arr {
		arr[i] = make([]byte, len(s))
	}

	row, col, isDiagnol := 0, 0, false
	for i := range s {
		arr[row][col] = s[i]
		if !isDiagnol {
			if row < numRows-1 {
				row++
			} else {
				row--
				col++
				isDiagnol = true
			}
		} else {
			if row == 0 {
				row++
				isDiagnol = false
			} else {
				row--
				col++
			}
		}
	}

	res := []byte{}
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] != 0 {
				res = append(res, arr[i][j])
			}
		}
	}
	return string(res)
}
