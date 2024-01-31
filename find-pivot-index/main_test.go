package findpivotindex

import (
	"fmt"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	test1 := pivotIndex([]int{1, 7, 3, 6, 5, 6})
	if test1 != 3 {
		fmt.Println("test1", test1)
		t.FailNow()
	}

	test2 := pivotIndex([]int{1, 2, 3})
	if test2 != -1 {
		fmt.Println("test2", test2)
		t.FailNow()
	}

	test3 := pivotIndex([]int{2, 1, -1})
	if test3 != 0 {
		fmt.Println("test3", test3)
		t.FailNow()
	}

	test4 := pivotIndex([]int{-1, -1, 0, 0, -1, -1})
	if test4 != 2 {
		fmt.Println("test4", test4)
		t.FailNow()
	}
}
