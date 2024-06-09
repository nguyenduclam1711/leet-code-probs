package wordpattern

import (
	"fmt"
	"testing"
)

func TestWordPattern(t *testing.T) {
	test1 := wordPattern("abba", "dog cat cat dog")
	if !test1 {
		fmt.Println("test1", test1)
		t.FailNow()
	}

	test2 := wordPattern("abba", "dog cat cat fish")
	if test2 {
		fmt.Println("test2", test2)
		t.FailNow()
	}

	test3 := wordPattern("aaaa", "dog cat cat dog")
	if test3 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
