package minimumwindowsubstring

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	test1 := minWindow("ADOBECODEBANC", "ABC")
	if test1 != "BANC" {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := minWindow("a", "a")
	if test2 != "a" {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := minWindow("a", "aa")
	if test3 != "" {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
