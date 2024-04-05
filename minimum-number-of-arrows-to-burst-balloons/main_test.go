package minimumnumberofarrowstoburstballoons

import (
	"fmt"
	"testing"
)

func TestFindMinArrowShots(t *testing.T) {
	test1 := findMinArrowShots([][]int{
		{10, 16},
		{2, 8},
		{1, 6},
		{7, 12},
	})
	if test1 != 2 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := findMinArrowShots([][]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	})
	if test2 != 4 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := findMinArrowShots([][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	})
	if test3 != 2 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := findMinArrowShots([][]int{
		{9, 12},
		{1, 10},
		{4, 11},
		{8, 12},
		{3, 9},
		{6, 9},
		{6, 7},
	})
	if test4 != 2 {
		fmt.Println("test4", test4)
		t.FailNow()
	}
}
