package rotatearray

func rotate(nums []int, k int) {
	newS := make([]int, len(nums))
	for i, n := range nums {
		newPos := (i + k) % len(nums)
		newS[newPos] = n
	}
	copy(nums, newS)
}
