package twosum

func TwoSum(nums []int, target int) []int {
	mapping := make(map[int]int)

	for i, num := range nums {
		if _, exist := mapping[target-num]; exist {
			return []int{mapping[target-num], i}
		}
		mapping[num] = i
	}
	return []int{}
}
