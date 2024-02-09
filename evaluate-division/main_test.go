package evaluatedivision

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	test1 := calcEquation(
		[][]string{
			{"a", "b"},
			{"b", "c"},
		},
		[]float64{2, 3},
		[][]string{
			{"a", "c"},
			{"b", "a"},
			{"a", "e"},
			{"a", "a"},
			{"x", "x"},
		},
	)
	if !reflect.DeepEqual(test1, []float64{6, 0.5, -1, 1, -1}) {
		fmt.Println("test1", test1, len(test1))
		t.FailNow()
	}
	test2 := calcEquation(
		[][]string{
			{"a", "b"},
			{"b", "c"},
			{"bc", "cd"},
		},
		[]float64{1.5, 2.5, 5.0},
		[][]string{
			{"a", "c"},
			{"c", "b"},
			{"bc", "cd"},
			{"cd", "bc"},
		},
	)
	if !reflect.DeepEqual(test2, []float64{3.75, 0.4, 5, 0.2}) {
		fmt.Println("test2", test2, len(test2))
		t.FailNow()
	}
	test3 := calcEquation(
		[][]string{
			{"a", "b"},
		},
		[]float64{0.5},
		[][]string{
			{"a", "b"},
			{"b", "a"},
			{"a", "c"},
			{"x", "y"},
		},
	)
	if !reflect.DeepEqual(test3, []float64{0.5, 2, -1, -1}) {
		fmt.Println("test3", test3, len(test3))
		t.FailNow()
	}
	test4 := calcEquation(
		[][]string{
			{"a", "b"},
			{"c", "d"},
			{"e", "f"},
			{"g", "h"},
		},
		[]float64{4.5, 2.3, 8.9, 0.44},
		[][]string{
			{"a", "c"},
			{"d", "f"},
			{"h", "e"},
			{"b", "e"},
			{"d", "h"},
			{"g", "f"},
			{"c", "g"},
		},
	)
	if !reflect.DeepEqual(test4, []float64{-1, -1, -1, -1, -1, -1, -1}) {
		fmt.Println("test4", test4, len(test4))
		t.FailNow()
	}
}
