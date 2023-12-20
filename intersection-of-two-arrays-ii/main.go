package intersectionoftwoarraysii

func Intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	uniqNums := map[int]int{}

	for _, num := range nums1 {
		uniqNums[num]++
	}
	for _, num := range nums2 {
		if uniqNums[num] > 0 {
			result = append(result, num)
			uniqNums[num]--
		}
	}
	return result
}
