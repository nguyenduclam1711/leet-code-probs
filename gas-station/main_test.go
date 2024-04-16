package gasstation

import (
	"fmt"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	test1 := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	if test1 != 3 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})
	if test2 != -1 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
