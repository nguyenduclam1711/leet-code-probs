package summaryranges

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	low, hi := nums[0], nums[0]
	res := []string{}
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		if n == hi+1 {
			hi = n
			continue
		}
		if low == hi {
			res = append(res, fmt.Sprint(low))
		} else {
			res = append(res, fmt.Sprint(low, "->", hi))
		}
		low = n
		hi = n
	}
	if low == hi {
		res = append(res, fmt.Sprint(low))
	} else {
		res = append(res, fmt.Sprint(low, "->", hi))
	}
	return res
}
