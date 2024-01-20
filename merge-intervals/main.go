package mergeintervals

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	result := [][]int{
		intervals[0],
	}
	i, j := 0, 1

	for j < len(intervals) {
		resInterval := result[i]
		currInterval := intervals[j]

		if currInterval[0] > resInterval[1] {
			result = append(result, currInterval)
			i++
		} else if currInterval[1] > resInterval[1] {
			resInterval[1] = currInterval[1]
		}
		j++
	}
	return result
}
