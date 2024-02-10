package kthlargestelementinanarray

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	test1 := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	if test1 != 5 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
	if test2 != 4 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
