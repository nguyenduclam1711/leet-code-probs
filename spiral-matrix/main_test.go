package spiralmatrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	test1 := spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})

	if !reflect.DeepEqual([]int{1, 2, 3, 6, 9, 8, 7, 4, 5}, test1) {
		t.Fail()
		fmt.Println("test1", test1)
	}

	test2 := spiralOrder([][]int{{2, 5, 8}, {4, 0, -1}})

	if !reflect.DeepEqual([]int{2, 5, 8, -1, 0, 4}, test2) {
		t.Fail()
		fmt.Println("test2", test2)
	}
}
