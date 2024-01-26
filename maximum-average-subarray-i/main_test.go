package maximumaveragesubarrayi

import (
	"fmt"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	test1 := findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)

	if test1 != 12.75 {
		t.Fail()
		fmt.Println("test1", test1)
	}

	test2 := findMaxAverage([]int{-1}, 1)

	if test2 != -1.0 {
		t.Fail()
		fmt.Println("test2", test2)
	}
}
