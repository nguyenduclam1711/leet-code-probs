package besttimetobuyandsellstockwithtransactionfee

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	test1 := maxProfit([]int{1, 3, 2, 8, 4, 9}, 2)
	if test1 != 8 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := maxProfit([]int{1, 3, 7, 5, 10, 3}, 3)
	if test2 != 6 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
