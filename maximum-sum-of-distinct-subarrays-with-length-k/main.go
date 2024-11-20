package maximumsumofdistinctsubarrayswithlengthk

func maximumSubarraySum(nums []int, k int) int64 {
	l, r := 0, k-1
	res := 0
	currNums := map[int]int{}
	currSum := 0

	for i := l; i <= r; i++ {
		currNums[nums[i]]++
		currSum += nums[i]
	}
	if len(currNums) == k && currSum > res {
		res = currSum
	}

	for r < len(nums) {
		currSum -= nums[l]
		currNums[nums[l]]--
		if currNums[nums[l]] == 0 {
			delete(currNums, nums[l])
		}
		l++
		r++
		if r == len(nums) {
			break
		}
		currNums[nums[r]]++
		currSum += nums[r]
		if len(currNums) == k && currSum > res {
			res = currSum
		}
	}
	return int64(res)
}
