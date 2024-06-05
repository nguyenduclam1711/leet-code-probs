package validsudoku

func isValid(m map[byte]bool, b byte) bool {
	if b == '.' {
		return true
	}
	if m[b] {
		return false
	}
	m[b] = true
	return true
}

func isRowValid(board [][]byte) bool {
	res := true
	for _, bytes := range board {
		m := map[byte]bool{}
		localRes := true
		for _, b := range bytes {
			valid := isValid(m, b)
			if !valid {
				localRes = false
				break
			}
		}
		if !localRes {
			res = false
			break
		}
	}
	return res
}

func isColValid(board [][]byte) bool {
	res := true

	for col := 0; col < 9; col++ {
		m := map[byte]bool{}
		localRes := true
		for row := 0; row < 9; row++ {
			valid := isValid(m, board[row][col])
			if !valid {
				localRes = false
				break
			}
		}
		if !localRes {
			res = false
			break
		}
	}
	return res
}

func isSmallBoardValid(board [][]byte) bool {
	res := true
	for loop := 0; loop < 9; loop++ {
		m := map[byte]bool{}
		localRes := true
		row := loop / 3 * 3
		maxRow := row + 3
		for row < maxRow {
			col := loop % 3 * 3
			maxCol := col + 3
			localRes2 := true
			for col < maxCol {
				valid := isValid(m, board[row][col])
				if !valid {
					localRes2 = false
					break
				}
				col++
			}
			if !localRes2 {
				localRes = false
				break
			}
			row++
		}
		if !localRes {
			res = false
			break
		}
	}
	return res
}

func isValidSudoku(board [][]byte) bool {
	return isRowValid(board) && isColValid(board) && isSmallBoardValid(board)
}
