package besttimetobuyandsellstockii

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	test1 := maxProfit([]int{7, 1, 5, 3, 6, 4})
	if test1 != 7 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := maxProfit([]int{1, 2, 3, 4, 5})
	if test2 != 4 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := maxProfit([]int{7, 6, 4, 3, 1})
	if test3 != 0 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
