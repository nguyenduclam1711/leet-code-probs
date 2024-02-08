package reorderroutestomakeallpathsleadtothecityzero

func dfs(curr int, cityCanLeadToZero []bool, mapConnections [][][]int, visited map[int]bool, res *int) {
	if visited[curr] {
		return
	}
	connections := mapConnections[curr]
	visited[curr] = true
	for _, connection := range connections {
		firstCity, secondCity := connection[0], connection[1]
		if cityCanLeadToZero[secondCity] {
			cityCanLeadToZero[firstCity] = true
			dfs(firstCity, cityCanLeadToZero, mapConnections, visited, res)
		} else if cityCanLeadToZero[firstCity] {
			*res++
			cityCanLeadToZero[secondCity] = true
			dfs(secondCity, cityCanLeadToZero, mapConnections, visited, res)
		}
	}
}

func minReorder(n int, connections [][]int) int {
	mapConnections := make([][][]int, n)
	for _, connection := range connections {
		firstCity, secondCity := connection[0], connection[1]
		mapConnections[firstCity] = append(mapConnections[firstCity], connection)
		mapConnections[secondCity] = append(mapConnections[secondCity], connection)
	}

	cityCanLeadToZero := make([]bool, n)
	cityCanLeadToZero[0] = true
	res := 0
	visited := map[int]bool{}
	for i := 0; i < n; i++ {
		dfs(i, cityCanLeadToZero, mapConnections, visited, &res)
	}
	return res
}
