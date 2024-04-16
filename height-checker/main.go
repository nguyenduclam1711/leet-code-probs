package heightchecker

import "sort"

func heightChecker(heights []int) int {
	copyHeights := make([]int, len(heights))
	copy(copyHeights, heights)
	sort.Ints(copyHeights)

	res := 0
	for i := range heights {
		if heights[i] != copyHeights[i] {
			res++
		}
	}
	return res
}
