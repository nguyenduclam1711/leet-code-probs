package mincostclimbingstairs

import (
	"fmt"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	test1 := minCostClimbingStairs([]int{10, 15, 20})
	if test1 != 15 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	if test2 != 6 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
