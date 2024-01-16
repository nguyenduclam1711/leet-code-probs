package numberofislands

func isNewIsland(x int, y int, grid [][]byte, visited [][]bool) bool {
	if x < 0 || x >= len(grid) {
		return false
	}
	if y < 0 || y >= len(grid[x]) {
		return false
	}
	if visited[x][y] {
		return false
	}
	visited[x][y] = true
	if grid[x][y] == '0' {
		return false
	}
	// move up
	up := isNewIsland(x-1, y, grid, visited)
	// move right
	right := isNewIsland(x, y+1, grid, visited)
	// move down
	down := isNewIsland(x+1, y, grid, visited)
	// move left
	left := isNewIsland(x, y-1, grid, visited)
	return !up || !right || !down || !left
}

func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	res := 0
	for i := range grid {
		for j := range grid[i] {
			if isNewIsland(i, j, grid, visited) {
				res++
			}
		}
	}
	return res
}
