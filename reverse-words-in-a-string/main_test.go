package reversewordsinastring

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	test1 := reverseWords("the sky is blue")

	if test1 != "blue is sky the" {
		t.Fail()
		fmt.Println("test1", test1)
	}

	test2 := reverseWords("  hello world  ")

	if test2 != "world hello" {
		t.Fail()
		fmt.Println("test2", test2)
	}

	test3 := reverseWords("a good   example")

	if test3 != "example good a" {
		t.Fail()
		fmt.Println("test3", test3)
	}
}
