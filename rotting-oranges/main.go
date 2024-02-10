package rottingoranges

func orangesRotting(grid [][]int) int {
	q := [][2]int{}
	visited := make([][]bool, len(grid))
	numberOfOranges := 0
	numberOfRottingOranges := 0
	// find rotting orange
	for i := range grid {
		visited[i] = make([]bool, len(grid[0]))
		for j := range grid[i] {
			if grid[i][j] != 0 {
				numberOfOranges++
			}
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
				visited[i][j] = true
				numberOfRottingOranges++
			}
		}
	}

	if numberOfOranges == 0 {
		return 0
	}

	addToQueue := func(x int, y int) {
		if x < 0 || x == len(grid) || y < 0 || y == len(grid[0]) {
			return
		}
		if grid[x][y] == 0 {
			return
		}
		if visited[x][y] {
			return
		}
		visited[x][y] = true
		if grid[x][y] == 1 {
			q = append(q, [2]int{x, y})
			numberOfRottingOranges++
		}
	}

	res := -1
	for len(q) > 0 {
		res++
		currLen := len(q)
		for _, pos := range q {
			x, y := pos[0], pos[1]
			addToQueue(x-1, y)
			addToQueue(x, y+1)
			addToQueue(x+1, y)
			addToQueue(x, y-1)
		}
		q = q[currLen:]
	}

	if numberOfRottingOranges != numberOfOranges {
		return -1
	}
	return res
}
