package dominoandtrominotiling

import (
	"fmt"
	"testing"
)

func TestNumTilings(t *testing.T) {
	test1 := numTilings(1)
	if test1 != 1 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := numTilings(2)
	if test2 != 2 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := numTilings(3)
	if test3 != 5 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := numTilings(4)
	if test4 != 11 {
		fmt.Println("test4", test4)
		t.FailNow()
	}
	test5 := numTilings(30)
	if test5 != 312342182 {
		fmt.Println("test5", test5)
		t.FailNow()
	}
}
