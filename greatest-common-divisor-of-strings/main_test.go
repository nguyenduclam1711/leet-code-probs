package greatestcommondivisorofstrings

import (
	"fmt"
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	test1 := gcdOfStrings("ABCABC", "ABC")

	if test1 != "ABC" {
		fmt.Println("test1", test1)
		t.Fail()
	}

	test2 := gcdOfStrings("ABABAB", "ABAB")

	if test2 != "AB" {
		fmt.Println("test2", test2)
		t.Fail()
	}

	test3 := gcdOfStrings("LEET", "CODE")

	if test3 != "" {
		fmt.Println("test3", test3)
		t.Fail()
	}
}
