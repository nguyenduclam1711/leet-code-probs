package findifarraycanbesorted

import "math"

func countSetBitsRecursive(n int) int {
	remainder := n % 2
	divideVal := n / 2
	res := 0

	if remainder == 1 {
		res++
	}
	if divideVal == 0 {
		return res
	}
	res += countSetBitsRecursive(divideVal)
	return res
}

func canSortArray(nums []int) bool {
	currSetBits := countSetBitsRecursive(nums[0])
	prevMax, currMin, currMax := math.MinInt, nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		n := nums[i]
		setBits := countSetBitsRecursive(n)
		if setBits == currSetBits {
			if n > currMax {
				currMax = n
			}
			if n < currMin {
				currMin = n
			}
		} else {
			prevMax = currMax
			currMin, currMax = n, n
			currSetBits = setBits
		}
		if currMin < prevMax {
			return false
		}
	}
	return true
}
