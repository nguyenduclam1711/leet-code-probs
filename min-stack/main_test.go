package minstack

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	test1 := Constructor()
	test1.Push(2)
	test1.Push(0)
	test1.Push(3)
	test1.Push(0)
	if test1.GetMin() != 0 {
		fmt.Println(test1.GetMin())
		t.FailNow()
	}
	test1.Pop()
	if test1.GetMin() != 0 {
		fmt.Println(test1.GetMin())
		t.FailNow()
	}
	test1.Pop()
	if test1.GetMin() != 0 {
		fmt.Println(test1.GetMin())
		t.FailNow()
	}
	test1.Pop()
	if test1.GetMin() != 2 {
		fmt.Println(test1.GetMin())
		t.FailNow()
	}
}
