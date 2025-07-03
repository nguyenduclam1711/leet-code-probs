package zigzagconversion

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	test1 := convert("PAYPALISHIRING", 3)
	if test1 != "PAHNAPLSIIGYIR" {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := convert("PAYPALISHIRING", 4)
	if test2 != "PINALSIGYAHRPI" {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := convert("AB", 1)
	if test3 != "AB" {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := convert("ABC", 1)
	if test4 != "ABC" {
		fmt.Println("test4", test4)
		t.FailNow()
	}
}
