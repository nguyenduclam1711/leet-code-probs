package reversevowelsofastring

import (
	"fmt"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	test1 := reverseVowels("leetcode")

	if test1 != "leotcede" {
		fmt.Println("test1", test1)
		t.Fail()
	}

	test2 := reverseVowels("hello")

	if test2 != "holle" {
		fmt.Println("test2", test2)
		t.Fail()
	}

	test3 := reverseVowels(" ")

	if test3 != " " {
		fmt.Println("test3", test3)
		t.Fail()
	}
}
