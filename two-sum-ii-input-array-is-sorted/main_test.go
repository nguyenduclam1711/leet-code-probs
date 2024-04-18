package twosumiiinputarrayissorted

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	test1 := twoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual([]int{1, 2}, test1) {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := twoSum([]int{2, 3, 4}, 6)
	if !reflect.DeepEqual([]int{1, 3}, test2) {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := twoSum([]int{-1, 0}, -1)
	if !reflect.DeepEqual([]int{1, 2}, test3) {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := twoSum([]int{5, 25, 75}, 100)
	if !reflect.DeepEqual([]int{2, 3}, test4) {
		fmt.Println("test4", test4)
		t.FailNow()
	}
}
