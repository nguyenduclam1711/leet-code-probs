package minimumflipstomakeaorbequaltoc

import (
	"fmt"
	"testing"
)

func TestMinFlips(t *testing.T) {
	test1 := minFlips(2, 6, 5)
	if test1 != 3 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := minFlips(4, 2, 7)
	if test2 != 1 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := minFlips(1, 2, 3)
	if test3 != 0 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := minFlips(8, 3, 5)
	if test4 != 3 {
		fmt.Println("test4", test4)
		t.FailNow()
	}
	test5 := minFlips(5, 2, 8)
	if test5 != 4 {
		fmt.Println("test5", test5)
		t.FailNow()
	}
}
