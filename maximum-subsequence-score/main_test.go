package maximumsubsequencescore

import (
	"fmt"
	"testing"
)

func TestMaxScore(t *testing.T) {
	test1 := maxScore(
		[]int{1, 3, 3, 2},
		[]int{2, 1, 3, 4},
		3,
	)
	if test1 != 12 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := maxScore(
		[]int{4, 2, 3, 1, 1},
		[]int{7, 5, 10, 9, 6},
		1,
	)
	if test2 != 30 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := maxScore(
		[]int{2, 1, 14, 12},
		[]int{11, 7, 13, 6},
		3,
	)
	if test3 != 168 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
