package movezeroes

func MoveZeroes(nums []int) {
	p := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
}
