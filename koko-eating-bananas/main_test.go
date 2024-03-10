package kokoeatingbananas

import (
	"fmt"
	"testing"
)

func TestMinEatingSpeed(t *testing.T) {
	test1 := minEatingSpeed([]int{3, 6, 7, 11}, 8)
	if test1 != 4 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := minEatingSpeed([]int{30, 11, 23, 4, 20}, 5)
	if test2 != 30 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := minEatingSpeed([]int{30, 11, 23, 4, 20}, 6)
	if test3 != 23 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := minEatingSpeed([]int{312884470}, 312884469)
	if test4 != 2 {
		fmt.Println("test4", test4)
		t.FailNow()
	}
	test5 := minEatingSpeed([]int{312884470}, 968709470)
	if test5 != 1 {
		fmt.Println("test5", test5)
		t.FailNow()
	}
	test6 := minEatingSpeed([]int{1000000000, 1000000000}, 3)
	if test6 != 1000000000 {
		fmt.Println("test6", test6)
		t.FailNow()
	}
}
