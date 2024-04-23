package candy

import (
	"fmt"
	"testing"
)

func TestCandy(t *testing.T) {
	test1 := candy([]int{1, 0, 2})
	if test1 != 5 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := candy([]int{1, 2, 2})
	if test2 != 4 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
