package totalcosttohirekworkers

import (
	"fmt"
	"testing"
)

func TestTotalCost(t *testing.T) {
	test1 := totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4)
	if test1 != 11 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := totalCost([]int{1, 2, 4, 1}, 3, 3)
	if test2 != 4 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := totalCost([]int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}, 11, 2)
	if test3 != 423 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := totalCost([]int{2, 2, 2, 2, 2, 2, 1, 4, 5, 5, 5, 5, 5, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 7, 3)
	if test4 != 13 {
		fmt.Println("test4", test4)
		t.FailNow()
	}
}
