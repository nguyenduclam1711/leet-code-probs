package evaluatereversepolishnotation

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	test1 := evalRPN([]string{"2", "1", "+", "3", "*"})
	if test1 != 9 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := evalRPN([]string{"4", "13", "5", "/", "+"})
	if test2 != 6 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	if test3 != 22 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
