package findthedifferenceoftwoarrays

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := map[int]bool{}
	map2 := map[int]bool{}
	for _, n := range nums1 {
		if !map1[n] {
			map1[n] = true
		}
	}
	for _, n := range nums2 {
		if !map2[n] {
			map2[n] = true
		}
	}

	res := make([][]int, 2)
	for k := range map1 {
		if !map2[k] {
			res[0] = append(res[0], k)
		}
	}
	for k := range map2 {
		if !map1[k] {
			res[1] = append(res[1], k)
		}
	}
	return res
}
