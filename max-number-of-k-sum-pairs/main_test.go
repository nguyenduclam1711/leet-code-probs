package maxnumberofksumpairs

import (
	"fmt"
	"testing"
)

func TestMaxOperations(t *testing.T) {
	test1 := maxOperations([]int{1, 2, 3, 4}, 5)

	if test1 != 2 {
		t.Fail()
		fmt.Println("test1", test1)
	}

	test2 := maxOperations([]int{3, 1, 3, 4, 3}, 6)

	if test2 != 1 {
		t.Fail()
		fmt.Println("test2", test2)
	}
}
