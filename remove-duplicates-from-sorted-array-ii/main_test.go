package removeduplicatesfromsortedarrayii

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	test1 := removeDuplicates([]int{1, 1, 1, 2, 2, 3})
	if test1 != 5 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
	if test2 != 7 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
