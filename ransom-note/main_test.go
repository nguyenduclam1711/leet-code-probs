package ransomnote

import (
	"fmt"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	test1 := canConstruct("a", "b")
	if test1 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := canConstruct("aa", "ab")
	if test2 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := canConstruct("aa", "aab")
	if !test3 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
