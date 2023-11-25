package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	uniqNums := map[int]int{}

	for _, num := range nums1 {
		uniqNums[num]++
	}
	for _, num := range nums2 {
		if uniqNums[num] > 0 {
			result = append(result, num)
			uniqNums[num]--
		}
	}
	return result
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1, 3, 3, 4, 5, 1}, []int{2, 2, 1, 1, 1, 3, 4, 5}))
}
