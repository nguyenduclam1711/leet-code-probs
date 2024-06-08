package isomorphicstrings

import (
	"fmt"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	test1 := isIsomorphic("egg", "add")
	if !test1 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := isIsomorphic("foo", "bar")
	if test2 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := isIsomorphic("paper", "title")
	if !test3 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
