package gameoflife

func checkNeighbor(i, j int, board [][]int, liveNeighbors, deadNeighbors *int) {
	// out of range
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
		*deadNeighbors++
		return
	}
	curr := board[i][j]
	if curr == 0 {
		*deadNeighbors++
	} else {
		*liveNeighbors++
	}
}

func checkAllNeighbors(i, j int, board [][]int, liveNeighbors, deadNeighbors *int) {
	checkNeighbor(i-1, j-1, board, liveNeighbors, deadNeighbors) // top left
	checkNeighbor(i-1, j, board, liveNeighbors, deadNeighbors)   // top
	checkNeighbor(i-1, j+1, board, liveNeighbors, deadNeighbors) // top right
	checkNeighbor(i, j+1, board, liveNeighbors, deadNeighbors)   // rigth
	checkNeighbor(i+1, j+1, board, liveNeighbors, deadNeighbors) // bottom right
	checkNeighbor(i+1, j, board, liveNeighbors, deadNeighbors)   // bottom
	checkNeighbor(i+1, j-1, board, liveNeighbors, deadNeighbors) // bottom left
	checkNeighbor(i, j-1, board, liveNeighbors, deadNeighbors)   // left
}

func checkRule(i, j int, board [][]int, newBoard [][]int) {
	curr := board[i][j]
	liveNeighbors, deadNeighbors := 0, 0

	checkAllNeighbors(i, j, board, &liveNeighbors, &deadNeighbors)
	if curr == 0 { // dead cell
		if liveNeighbors == 3 {
			newBoard[i][j] = 1
		}
	} else { // live cell
		if liveNeighbors == 2 || liveNeighbors == 3 {
			newBoard[i][j] = 1
		}
	}
}

func gameOfLife(board [][]int) {
	// init new board for replacing
	newBoard := make([][]int, len(board))
	for i := range newBoard {
		newBoard[i] = make([]int, len(board[i]))
	}

	for i := range board {
		for j := range board[i] {
			checkRule(i, j, board, newBoard)
		}
	}
	for i := range newBoard {
		copy(board[i], newBoard[i])
	}
}
