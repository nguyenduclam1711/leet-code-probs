package equalrowandcolumnpairs

import (
	"fmt"
	"testing"
)

func TestEqualPairs(t *testing.T) {
	test1 := equalPairs([][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7},
	})
	if test1 != 1 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := equalPairs([][]int{
		{3, 1, 2, 2},
		{1, 4, 4, 5},
		{2, 4, 2, 2},
		{2, 4, 2, 2},
	})
	if test2 != 3 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := equalPairs([][]int{
		{3, 1, 2, 2},
		{1, 4, 4, 4},
		{2, 4, 2, 2},
		{2, 5, 2, 2},
	})
	if test3 != 3 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := equalPairs([][]int{
		{11, 1},
		{1, 11},
	})
	if test4 != 2 {
		fmt.Println("test4", test4)
		t.FailNow()
	}
}
