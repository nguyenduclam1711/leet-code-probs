package courseschedule

func dfs(curr int, graph [][]int, visited map[int]bool, nodeCanBeFinish map[int]bool) bool {
	if nodeCanBeFinish[curr] {
		return true
	}
	if visited[curr] {
		return false
	}
	visited[curr] = true
	for _, n := range graph[curr] {
		if !dfs(n, graph, visited, nodeCanBeFinish) {
			return false
		}
	}
	visited[curr] = false
	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)

	for _, p := range prerequisites {
		graph[p[0]] = append(graph[p[0]], p[1])
	}
	visited := map[int]bool{}
	nodeCanBeFinish := map[int]bool{}
	for i := range graph {
		if !dfs(i, graph, visited, nodeCanBeFinish) {
			return false
		}
		nodeCanBeFinish[i] = true
	}
	return true
}
