package trappingrainwater

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	test1 := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	if test1 != 6 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := trap([]int{4, 2, 0, 3, 2, 5})
	if test2 != 9 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
