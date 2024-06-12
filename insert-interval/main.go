package insertinterval

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{
			newInterval,
		}
	}
	res := [][]int{}
	lo, hi := newInterval[0], newInterval[1]
	for _, interval := range intervals {
		if interval[0] <= hi && interval[1] >= lo {
			if interval[0] < lo {
				lo = interval[0]
			}
			if interval[1] > hi {
				hi = interval[1]
			}
		} else {
			if interval[0] > hi {
				res = append(res, []int{
					lo,
					hi,
				})
				lo, hi = interval[0], interval[1]
			} else if interval[1] < lo {
				res = append(res, interval)
			}
		}
	}
	res = append(res, []int{
		lo,
		hi,
	})
	return res
}
