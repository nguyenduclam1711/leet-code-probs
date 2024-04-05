package minimumnumberofarrowstoburstballoons

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	res := 1
	curMaxPoint := points[0][1]

	for i := 1; i < len(points); i++ {
		x, y := points[i][0], points[i][1]
		if x <= curMaxPoint {
			continue
		}
		curMaxPoint = y
		res++
	}
	return res
}
