package integertoroman

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	test1 := intToRoman(3)
	if test1 != "III" {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := intToRoman(58)
	if test2 != "LVIII" {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := intToRoman(1994)
	if test3 != "MCMXCIV" {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
