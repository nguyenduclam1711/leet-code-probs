package nearestexitfromentranceinmaze

import (
	"fmt"
	"testing"
)

func TestNearestExit(t *testing.T) {
	test1 := nearestExit(
		[][]byte{
			{'+', '+', '.', '+'},
			{'.', '.', '.', '+'},
			{'+', '+', '+', '.'},
		},
		[]int{1, 2},
	)
	if test1 != 1 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := nearestExit(
		[][]byte{
			{'+', '+', '+'},
			{'.', '.', '.'},
			{'+', '+', '+'},
		},
		[]int{1, 0},
	)
	if test2 != 2 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := nearestExit(
		[][]byte{
			{'.', '+'},
		},
		[]int{0, 0},
	)
	if test3 != -1 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4Data := [][]string{
		{"+", ".", "+", "+", "+", "+", "+"},
		{"+", ".", "+", ".", ".", ".", "+"},
		{"+", ".", "+", ".", "+", ".", "+"},
		{"+", ".", ".", ".", "+", ".", "+"},
		{"+", "+", "+", "+", "+", ".", "+"},
	}
	test4RealData := make([][]byte, len(test4Data))
	for i := range test4Data {
		test4RealData[i] = make([]byte, len(test4Data[i]))
		for j := range test4Data[i] {
			test4RealData[i][j] = test4Data[i][j][0]
		}
	}
	test4 := nearestExit(
		test4RealData,
		[]int{0, 1},
	)
	if test4 != 12 {
		fmt.Println("test4", test4)
		t.FailNow()
	}
	test5Data := [][]string{
		{"+", ".", "+", "+", "+", "+", "+"},
		{"+", ".", "+", ".", ".", ".", "+"},
		{"+", ".", "+", ".", "+", ".", "+"},
		{"+", ".", ".", ".", ".", ".", "+"},
		{"+", "+", "+", "+", ".", "+", "."},
	}
	test5RealData := make([][]byte, len(test5Data))
	for i := range test5Data {
		test5RealData[i] = make([]byte, len(test5Data[i]))
		for j := range test5Data[i] {
			test5RealData[i][j] = test5Data[i][j][0]
		}
	}
	test5 := nearestExit(
		test5RealData,
		[]int{0, 1},
	)
	if test5 != 7 {
		fmt.Println("test5", test5)
		t.FailNow()
	}
}
