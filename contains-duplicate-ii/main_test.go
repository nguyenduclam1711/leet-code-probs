package containsduplicateii

import (
	"fmt"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	test1 := containsNearbyDuplicate([]int{1, 2, 3, 1}, 3)
	if !test1 {
		fmt.Println("test1", test1)
		t.FailNow()
	}

	test2 := containsNearbyDuplicate([]int{1, 0, 1, 1}, 1)
	if !test2 {
		fmt.Println("test2", test2)
		t.FailNow()
	}

	test3 := containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)
	if test3 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
