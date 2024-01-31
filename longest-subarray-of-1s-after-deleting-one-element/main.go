package longestsubarrayof1safterdeletingoneelement

func longestSubarray(nums []int) int {
	l, r := 0, 0
	numberOfZeroes := 0
	res := 0

	for r < len(nums) {
		if nums[r] == 0 {
			numberOfZeroes++
		}
		if numberOfZeroes <= 1 {
			windowLen := r - l
			if windowLen > res {
				res = windowLen
			}
		} else {
			if nums[l] == 0 {
				numberOfZeroes--
			}
			l++
		}
		r++
	}
	return res
}
