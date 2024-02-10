package kthlargestelementinanarray

import (
	"container/heap"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push appends an element to the heap.
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the maximum element from the heap.
func (h *MaxHeap) Pop() interface{} {
	currLen := len(*h)
	x := (*h)[currLen-1]
	*h = (*h)[0 : currLen-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)
	for _, n := range nums {
		heap.Push(h, n)
	}
	var curr int
	for i := 1; i <= k; i++ {
		curr = heap.Pop(h).(int)
	}
	return curr
}
