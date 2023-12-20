package intersectionoftwoarrays

func Intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	uniqNums := map[int]bool{}

	for _, num := range nums1 {
		if !uniqNums[num] {
			uniqNums[num] = true
		}
	}
	for _, num := range nums2 {
		if uniqNums[num] {
			result = append(result, num)
			delete(uniqNums, num)
		}
	}
	return result
}
