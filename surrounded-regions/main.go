package surroundedregions

const (
	NOT_VISITED    = 0
	VISITED        = 1
	NOT_SURROUNDED = 2
)

func markNotSurroundedRegions(x, y int, board [][]byte, m, n int, visited [][]int) {
	if x < 0 || x >= m || y < 0 || y >= n {
		return
	}
	if visited[x][y] == NOT_SURROUNDED || visited[x][y] == VISITED {
		return
	}
	if board[x][y] == 'X' {
		visited[x][y] = VISITED
		return
	}
	visited[x][y] = NOT_SURROUNDED
	markNotSurroundedRegions(x-1, y, board, m, n, visited) // top
	markNotSurroundedRegions(x, y+1, board, m, n, visited) // right
	markNotSurroundedRegions(x+1, y, board, m, n, visited) // bottom
	markNotSurroundedRegions(x, y-1, board, m, n, visited) // top
}

func replaceAllSurroundedRegions(x, y int, board [][]byte, m, n int, visited [][]int) {
	if x < 0 || x >= m || y < 0 || y >= n {
		return
	}
	if visited[x][y] == NOT_SURROUNDED || visited[x][y] == VISITED {
		return
	}
	visited[x][y] = VISITED
	if board[x][y] == 'X' {
		return
	}
	board[x][y] = 'X'
	replaceAllSurroundedRegions(x-1, y, board, m, n, visited) // top
	replaceAllSurroundedRegions(x, y+1, board, m, n, visited) // right
	replaceAllSurroundedRegions(x+1, y, board, m, n, visited) // bottom
	replaceAllSurroundedRegions(x, y-1, board, m, n, visited) // top
}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}

	// find all edge region that has been marked 'O'
	for col := range board[0] {
		markNotSurroundedRegions(0, col, board, m, n, visited)
		markNotSurroundedRegions(m-1, col, board, m, n, visited)
	}
	for row := range board {
		markNotSurroundedRegions(row, 0, board, m, n, visited)
		markNotSurroundedRegions(row, n-1, board, m, n, visited)
	}
	// replace all regions inside but not start from the edges
	for x := 1; x < m-1; x++ {
		for y := 1; y < n-1; y++ {
			replaceAllSurroundedRegions(x, y, board, m, n, visited)
		}
	}
}
