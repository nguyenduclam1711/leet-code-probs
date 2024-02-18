package findpeakelement

import (
	"fmt"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	test1 := findPeakElement([]int{1, 2, 3, 1})
	if test1 != 2 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})
	if test2 != 5 && test2 != 2 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := findPeakElement([]int{2, 1})
	if test3 != 0 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
