package simplifypath

import (
	"fmt"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	test1 := simplifyPath("/home//foo/")
	if test1 != "/home/foo" {
		fmt.Println("test1", test1)
		t.FailNow()
	}
}
