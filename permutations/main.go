package permutations

func permute(nums []int) [][]int {
	visitedPos := map[int]bool{}
	res := [][]int{}
	s := []int{}

	var recurse func(pos int)
	recurse = func(pos int) {
		if visitedPos[pos] {
			return
		}
		visitedPos[pos] = true
		s = append(s, nums[pos])

		if len(s) == len(nums) {
			newPerm := make([]int, len(nums))
			copy(newPerm, s)
			res = append(res, newPerm)
		} else {
			for i := range nums {
				recurse(i)
			}
		}
		visitedPos[pos] = false
		s = s[:len(s)-1]
	}

	for i := range nums {
		recurse(i)
	}
	return res
}
