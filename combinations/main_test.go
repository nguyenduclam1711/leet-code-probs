package combinations

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	test1 := combine(3, 2)
	if !reflect.DeepEqual(test1, [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}) {
		fmt.Println("test1", test1)
		t.Fail()
	}
}
