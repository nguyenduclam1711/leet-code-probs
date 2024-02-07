package keysandrooms

func CanVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	totalVisited := 0
	var dfs func(roomNumber int) bool
	dfs = func(roomNumber int) bool {
		if visited[roomNumber] {
			return false
		}
		visited[roomNumber] = true
		totalVisited++

		if totalVisited == len(rooms) {
			return true
		}
		for _, key := range rooms[roomNumber] {
			if dfs(key) {
				return true
			}
		}
		return false
	}
	return dfs(0)
}
