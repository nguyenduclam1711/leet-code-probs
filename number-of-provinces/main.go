package numberofprovinces

func dfs(city int, isConnected [][]int, visited []bool) {
	if visited[city] {
		return
	}
	visited[city] = true
	for nextCity, canBeConnect := range isConnected[city] {
		if canBeConnect == 1 {
			dfs(nextCity, isConnected, visited)
		}
	}
}

func FindCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	res := 0
	for city := range isConnected {
		if !visited[city] {
			dfs(city, isConnected, visited)
			res++
		}
	}
	return res
}
