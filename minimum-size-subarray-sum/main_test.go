package minimumsizesubarraysum

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	test1 := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	if test1 != 2 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := minSubArrayLen(4, []int{1, 4, 4})
	if test2 != 1 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})
	if test3 != 0 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := minSubArrayLen(15, []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8})
	if test4 != 2 {
		fmt.Println("test4", test4)
		t.FailNow()
	}
	test5 := minSubArrayLen(7, []int{1, 1, 1, 1, 7})
	if test5 != 1 {
		fmt.Println("test5", test5)
		t.FailNow()
	}
}
