package findthehighestaltitude

func largestAltitude(gain []int) int {
	curr := 0
	res := 0

	for _, n := range gain {
		curr += n
		if curr > res {
			res = curr
		}
	}
	return res
}
