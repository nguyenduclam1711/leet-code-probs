package setmatrixzeroes

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	matrix1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix1)
	if !reflect.DeepEqual([][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, matrix1) {
		t.Fail()
	}
}
