package rotatearray

func rotate(nums []int, k int) {
	newS := make([]int, len(nums))
	for i, n := range nums {
		newPos := (i + k) % len(nums)
		newS[newPos] = n
	}
	copy(nums, newS)
}

// O(1) extra space here
func rotate2(nums []int, k int) {
	k = k % len(nums)
	reverse := func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	// first reverse the entire arr
	reverse(0, len(nums)-1)
	// then reverse the first part
	reverse(0, k-1)
	// reverse the second part
	reverse(k, len(nums)-1)
}
