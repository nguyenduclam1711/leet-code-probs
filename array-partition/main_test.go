package arraypartition

import (
	"fmt"
	"testing"
)

func TestArrayPairSum(t *testing.T) {
	test1 := []int{1, 4, 3, 2}
	if arrayPairSum(test1) != 4 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := []int{6, 2, 6, 5, 1, 2}
	if arrayPairSum(test2) != 9 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
