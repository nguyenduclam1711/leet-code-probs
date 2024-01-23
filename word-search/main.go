package wordsearch

func dfs(x int, y int, board [][]byte, word string, visited [][]bool) bool {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return false
	}
	if visited[x][y] {
		return false
	}
	visited[x][y] = true
	var isValid bool

	firstChar := word[0]
	if board[x][y] != firstChar {
		isValid = false
	} else {
		nextWord := word[1:]

		if len(nextWord) == 0 {
			isValid = true
		} else {
			up := dfs(x-1, y, board, nextWord, visited)
			right := dfs(x, y+1, board, nextWord, visited)
			down := dfs(x+1, y, board, nextWord, visited)
			left := dfs(x, y-1, board, nextWord, visited)
			isValid = up || right || down || left
		}
	}
	visited[x][y] = false
	return isValid
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}

	for i := range board {
		for j := range board[i] {
			checkIfExist := dfs(i, j, board, word, visited)
			if checkIfExist {
				return true
			}
		}
	}
	return false
}
