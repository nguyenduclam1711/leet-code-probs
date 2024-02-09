package evaluatedivision

func dfs(start string, des string, mapValues map[string]map[string]float64, visited map[string]bool) float64 {
	if visited[start] {
		return -1
	}
	visited[start] = true
	var result float64 = -1
	if mapValues[start] == nil {
		mapValues[start] = map[string]float64{}
	}
	if mapValues[start][des] > 0 {
		result = mapValues[start][des]
	} else {
		for nextStr, valToNextStr := range mapValues[start] {
			if valToNextStr <= 0 {
				continue
			}
			valNextStrToDes := dfs(nextStr, des, mapValues, visited)
			if valNextStrToDes > 0 {
				result = valToNextStr * valNextStrToDes
				break
			}
		}
	}
	visited[start] = false
	mapValues[start][des] = result
	return result
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	mapValues := map[string]map[string]float64{}
	for i, equation := range equations {
		str, str2 := equation[0], equation[1]
		if mapValues[str] == nil {
			mapValues[str] = map[string]float64{}
		}
		if mapValues[str2] == nil {
			mapValues[str2] = map[string]float64{}
		}
		mapValues[str][str2] = values[i]
		mapValues[str][str] = 1
		mapValues[str2][str] = 1 / values[i]
		mapValues[str2][str2] = 1
	}

	result := []float64{}
	visited := map[string]bool{}
	for _, query := range queries {
		start, des := query[0], query[1]
		val := dfs(start, des, mapValues, visited)
		result = append(result, val)
	}
	return result
}
