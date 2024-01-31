package findthedifferenceoftwoarrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindDifference(t *testing.T) {
	test1 := findDifference([]int{1, 2, 3}, []int{2, 4, 6})
	if !reflect.DeepEqual([]int{1, 3}, test1[0]) && !reflect.DeepEqual([]int{4, 6}, test1[1]) {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2})
	if !reflect.DeepEqual([]int{3}, test2[0]) && len(test2) != 0 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
