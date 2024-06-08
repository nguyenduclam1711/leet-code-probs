package gameoflife

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGameOfLife(t *testing.T) {
	test1 := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(test1)
	if !reflect.DeepEqual([][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}, test1) {
		fmt.Println("test1", test1)
		t.FailNow()
	}

	test2 := [][]int{
		{1, 1},
		{1, 0},
	}
	gameOfLife(test2)
	if !reflect.DeepEqual([][]int{
		{1, 1},
		{1, 1},
	}, test2) {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
