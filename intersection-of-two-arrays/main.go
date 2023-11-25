package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	uniqNums := map[int]bool{}

	for _, num := range nums1 {
		if !uniqNums[num] {
			uniqNums[num] = true
		}
	}
	for _, num := range nums2 {
		if uniqNums[num] {
			result = append(result, num)
			delete(uniqNums, num)
		}
	}
	return result
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}
