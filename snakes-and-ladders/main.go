package snakesandladders

func getPos(squareNumber int, n int) (row, col int) {
	startRow := (squareNumber - 1) / n
	row = n - startRow - 1
	col = (squareNumber - 1) % n
	if startRow%2 != 0 {
		col = n - 1 - col
	}
	return row, col
}

func getNextSquares(currSquareNumber int) (minSquare, maxSquare int) {
	minSquare = currSquareNumber + 1
	maxSquare = currSquareNumber + 6
	return minSquare, maxSquare
}

func snakesAndLadders(board [][]int) int {
	q := []int{1}
	n := len(board)
	endSquare := n * n
	res := 0
	canReachToTheEnd := false
	visitedSquare := map[int]bool{}

	for len(q) > 0 {
		newQ := []int{}

		for _, squareNumber := range q {
			if visitedSquare[squareNumber] {
				continue
			}
			visitedSquare[squareNumber] = true
			if squareNumber == endSquare {
				canReachToTheEnd = true
				break
			}
			minSquare, maxSquare := getNextSquares(squareNumber)
			for nextSquare := minSquare; nextSquare <= maxSquare && nextSquare <= endSquare; nextSquare++ {
				row, col := getPos(nextSquare, n)
				if board[row][col] == -1 {
					newQ = append(newQ, nextSquare)
				} else {
					newQ = append(newQ, board[row][col])
				}
			}
		}
		if canReachToTheEnd {
			break
		}
		res++
		q = newQ
	}
	if canReachToTheEnd {
		return res
	}
	return -1
}
