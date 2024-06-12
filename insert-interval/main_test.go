package insertinterval

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	test1 := insert(
		[][]int{
			{1, 3},
			{6, 9},
		},
		[]int{2, 5},
	)
	if !reflect.DeepEqual(test1, [][]int{{1, 5}, {6, 9}}) {
		fmt.Println("test1", test1)
		t.FailNow()
	}
}
