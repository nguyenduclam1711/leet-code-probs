package nonoverlappingintervals

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 2 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := 0
	endInterval := intervals[0]

	for i := 1; i < len(intervals); i++ {
		currInterval := intervals[i]

		if currInterval[0] >= endInterval[1] {
			endInterval = currInterval
		} else {
			// meaning overlap, must remove 1 interval
			res++
			// take the interval that has earlier end => chance less likely to overlap the next one
			if currInterval[1] < endInterval[1] {
				endInterval = currInterval
			}
		}
	}
	return res
}
