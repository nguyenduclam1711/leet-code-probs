package reorderroutestomakeallpathsleadtothecityzero

import (
	"fmt"
	"testing"
)

func TestMinReorder(t *testing.T) {
	test1 := minReorder(6, [][]int{
		{0, 1},
		{1, 3},
		{2, 3},
		{4, 0},
		{4, 5},
	})
	if test1 != 3 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := minReorder(5, [][]int{
		{4, 3},
		{2, 3},
		{1, 2},
		{1, 0},
	})
	if test2 != 2 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := minReorder(6, [][]int{
		{0, 2},
		{0, 3},
		{4, 1},
		{4, 5},
		{5, 0},
	})
	if test3 != 3 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
