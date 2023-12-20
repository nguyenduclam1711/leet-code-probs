package majorityelement

func MajorityElement(nums []int) int {
	appearTimes := len(nums) / 2
	m := map[int]int{}

	for _, num := range nums {
		m[num]++
		if m[num] > appearTimes {
			return num
		}
	}
	return -1
}
