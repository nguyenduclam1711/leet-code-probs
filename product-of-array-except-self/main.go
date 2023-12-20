package productofarrayexceptself

func ProductExceptSelf(nums []int) []int {
	numsLen := len(nums)
	productPrefix := make([]int, numsLen)
	productPostfix := make([]int, numsLen)
	result := make([]int, numsLen)

	// calculate product prefix and postfix
	for i, num := range nums {
		if i == 0 {
			productPrefix[i] = num
			productPostfix[numsLen-1] = nums[numsLen-1]
		} else {
			productPrefix[i] = productPrefix[i-1] * num
			productPostfix[numsLen-1-i] = productPostfix[numsLen-i] * nums[numsLen-1-i]
		}
	}
	// calculate product result
	for j := range productPrefix {
		switch {
		case j == 0:
			{
				result[j] = productPostfix[j+1]
			}
		case j == numsLen-1:
			{
				result[j] = productPrefix[j-1]
			}
		default:
			{
				result[j] = productPrefix[j-1] * productPostfix[j+1]
			}
		}
	}
	return result
}
