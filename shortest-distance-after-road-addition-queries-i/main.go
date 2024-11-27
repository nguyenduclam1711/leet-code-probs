package shortestdistanceafterroadadditionqueriesi

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	mapGraph := make([][]int, n-1)

	for i := 0; i < n-1; i++ {
		mapGraph[i] = []int{i + 1}
	}

	res := []int{}

	for _, query := range queries {
		mapGraph[query[0]] = append(mapGraph[query[0]], query[1])
		count := 0
		queue := []int{0}
		visited := make([]bool, n)

		for len(queue) > 0 {
			newQueue := []int{}
			reachToEnd := false
			for _, currPoint := range queue {
				if visited[currPoint] {
					continue
				}
				visited[currPoint] = true
				if currPoint == n-1 {
					reachToEnd = true
					break
				}
				newQueue = append(newQueue, mapGraph[currPoint]...)
			}
			if reachToEnd {
				break
			}
			count++
			queue = newQueue
		}
		res = append(res, count)
	}
	return res
}
