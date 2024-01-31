package uniquenumberofoccurrences

import (
	"fmt"
	"testing"
)

func TestUniqueOccurrences(t *testing.T) {
	test1 := uniqueOccurrences([]int{1, 2, 2, 1, 1, 3})
	if !test1 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := uniqueOccurrences([]int{1, 2})
	if test2 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0})
	if !test3 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
