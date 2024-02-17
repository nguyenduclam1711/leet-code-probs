package maximumsubsequencescore

import (
	"container/heap"
	"sort"
)

type MinHeap [][]int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push appends an element to the heap.
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

// Pop removes and returns the maximum element from the heap.
func (h *MinHeap) Pop() interface{} {
	currLen := len(*h)
	x := (*h)[currLen-1]
	*h = (*h)[0 : currLen-1]
	return x
}

func findMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	arr := [][]int{}
	for i := 0; i < len(nums1); i++ {
		arr = append(arr, []int{nums1[i], nums2[i]})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})

	minHeap := &MinHeap{}
	sum := 0

	for i := 0; i < k; i++ {
		heap.Push(minHeap, arr[i])
		sum += arr[i][0]
	}

	res := sum * arr[k-1][1]

	for i := k; i < len(arr); i++ {
		smallNum := heap.Pop(minHeap).([]int)
		sum += arr[i][0] - smallNum[0]
		heap.Push(minHeap, arr[i])
		res = findMax(res, sum*arr[i][1])
	}
	return int64(res)
}
