package maxconsecutiveonesiii

func longestOnes(nums []int, k int) int {
	l, r := 0, 0
	numberOfZeroes := 0
	res := 0

	for r < len(nums) {
		if nums[r] == 0 {
			numberOfZeroes++
		}
		if numberOfZeroes > k {
			if nums[l] == 0 {
				numberOfZeroes--
			}
			l++
		}
		windowLen := r - l + 1
		if windowLen > res {
			res = windowLen
		}
		r++
	}
	return res
}
