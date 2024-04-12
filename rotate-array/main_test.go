package rotatearray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr1, 3)
	if !reflect.DeepEqual([]int{5, 6, 7, 1, 2, 3, 4}, arr1) {
		fmt.Println("arr1", arr1)
		t.FailNow()
	}
	arr2 := []int{-1, -100, 3, 99}
	rotate(arr2, 2)
	if !reflect.DeepEqual([]int{3, 99, -1, -100}, arr2) {
		fmt.Println("arr2", arr2)
		t.FailNow()
	}
	arr3 := []int{1, 2}
	rotate(arr3, 3)
	if !reflect.DeepEqual([]int{2, 1}, arr3) {
		fmt.Println("arr3", arr3)
		t.FailNow()
	}
}
