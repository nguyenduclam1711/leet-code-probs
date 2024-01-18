package pacificatlanticwaterflow

type Pos struct {
	visited      bool
	nearPacific  bool
	nearAtlantic bool
}

func isNearPacific(x int, y int) bool {
	return x == 0 || y == 0
}

func isNearAtlantic(x int, y int, heights [][]int) bool {
	return x == len(heights)-1 || y == len(heights[x])-1
}

func isCurrNodeValid(x int, y int, heights [][]int, prevX int, prevY int) bool {
	if prevX < 0 && prevY < 0 {
		return true
	}
	return x >= 0 && x < len(heights) && y >= 0 && y < len(heights[prevX]) && heights[x][y] <= heights[prevX][prevY]
}

func dfs(x int, y int, heights [][]int, positions [][]Pos, prevX int, prevY int) Pos {
	if !isCurrNodeValid(x, y, heights, prevX, prevY) {
		return Pos{}
	}
	if positions[x][y].visited || (positions[x][y].nearAtlantic && positions[x][y].nearPacific) {
		return positions[x][y]
	}
	positions[x][y].visited = true
	nearPacific := isNearPacific(x, y)
	nearAtlantic := isNearAtlantic(x, y, heights)
	// going up
	up := dfs(x-1, y, heights, positions, x, y)
	// going left
	left := dfs(x, y-1, heights, positions, x, y)
	// going right
	right := dfs(x, y+1, heights, positions, x, y)
	// going down
	down := dfs(x+1, y, heights, positions, x, y)

	canBeNearPacific := nearPacific || up.nearPacific || left.nearPacific || right.nearPacific || down.nearPacific
	canBeNearAtlantic := nearAtlantic || up.nearAtlantic || left.nearAtlantic || right.nearAtlantic || down.nearAtlantic

	if canBeNearPacific {
		positions[x][y].nearPacific = true
	}
	if canBeNearAtlantic {
		positions[x][y].nearAtlantic = true
	}
	positions[x][y].visited = false
	return positions[x][y]
}

func pacificAtlantic(heights [][]int) [][]int {
	positions := make([][]Pos, len(heights))
	for i := range positions {
		positions[i] = make([]Pos, len(heights[i]))
	}

	res := [][]int{}
	// traverse
	for i := range heights {
		for j := range heights[i] {
			pos := dfs(i, j, heights, positions, -1, -1)
			positions[i][j].visited = true
			if pos.nearAtlantic && pos.nearPacific {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
