package combinationsumiii

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	test1 := combinationSum3(3, 7)
	if !reflect.DeepEqual(test1, [][]int{{1, 2, 4}}) {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := combinationSum3(3, 9)
	if !reflect.DeepEqual(test2, [][]int{
		{1, 2, 6},
		{1, 3, 5},
		{2, 3, 4},
	}) {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := combinationSum3(4, 1)
	if !reflect.DeepEqual(test3, [][]int{}) {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
