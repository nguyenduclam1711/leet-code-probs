package main

import "fmt"

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	l, r, p := 0, len(nums), (0+len(nums))/2

	for l < r {
		if nums[p] < target {
			l = p + 1
		} else if nums[p] > target {
			r = p
		} else {
			// calculate here
			result[0] = p
			result[1] = p
			p1, p2 := p, p+1

			for p1 >= 0 {
				if nums[p1] == target {
					result[0] = p1
					p1--
				} else {
					break
				}
			}
			for p2 < len(nums) {
				if nums[p2] == target {
					result[1] = p2
					p2++
				} else {
					break
				}
			}
			break
		}
		p = (l + r) / 2
	}
	return result
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
