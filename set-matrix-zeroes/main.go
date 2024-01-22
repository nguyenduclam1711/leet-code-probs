package setmatrixzeroes

func setZeroes(matrix [][]int) {
	rows := []int{}
	columns := []int{}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rows = append(rows, i)
				columns = append(columns, j)
			}
		}
	}
	for _, row := range rows {
		for i := range matrix[row] {
			matrix[row][i] = 0
		}
	}
	for _, column := range columns {
		for i := range matrix {
			matrix[i][column] = 0
		}
	}
}
