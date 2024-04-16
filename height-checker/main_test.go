package heightchecker

import (
	"fmt"
	"testing"
)

func TestHeightChecker(t *testing.T) {
	test1 := []int{1, 1, 4, 2, 1, 3}
	if heightChecker(test1) != 3 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := []int{5, 1, 2, 3, 4}
	if heightChecker(test2) != 5 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := []int{1, 2, 3, 4, 5}
	if heightChecker(test3) != 0 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
