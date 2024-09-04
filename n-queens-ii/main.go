package nqueensii

func totalNQueens(n int) int {
	res := 0
	numberOfQueens := 0
	takenRows := map[int]bool{}
	takenCols := map[int]bool{}
	takenFirstDiags := map[int]bool{}
	takenSecondDiags := map[int]bool{}

	var recurse func(row int)
	recurse = func(row int) {
		if numberOfQueens == n {
			res++
			return
		}
		if row == n {
			return
		}
		if takenRows[row] {
			return
		}
		for col := 0; col < n; col++ {
			if takenCols[col] {
				continue
			}
			firstDiagIndex := row - col
			if takenFirstDiags[firstDiagIndex] {
				continue
			}
			secondDiagIndex := (n - row) - col
			if takenSecondDiags[secondDiagIndex] {
				continue
			}
			takenRows[row] = true
			takenCols[col] = true
			takenFirstDiags[firstDiagIndex] = true
			takenSecondDiags[secondDiagIndex] = true
			numberOfQueens++

			recurse(row + 1)

			numberOfQueens--
			takenRows[row] = false
			takenCols[col] = false
			takenFirstDiags[firstDiagIndex] = false
			takenSecondDiags[secondDiagIndex] = false
		}
	}
	recurse(0)
	return res
}
