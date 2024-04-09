package thekthfactorofn

import (
	"fmt"
	"testing"
)

func TestKthFactor(t *testing.T) {
	test1 := kthFactor(12, 3)
	if test1 != 3 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := kthFactor(7, 2)
	if test2 != 7 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := kthFactor(4, 4)
	if test3 != -1 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
