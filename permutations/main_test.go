package permutations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	test1 := permute([]int{1, 2, 3})
	if !reflect.DeepEqual(test1, [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}) {
		fmt.Println("test1", test1)
		t.Fail()
	}
	test2 := permute([]int{0, 1})
	if !reflect.DeepEqual(test2, [][]int{
		{0, 1},
		{1, 0},
	}) {
		fmt.Println("test2", test2)
		t.Fail()
	}
}
