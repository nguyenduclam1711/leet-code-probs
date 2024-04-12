package jumpgameii

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	test1 := jump([]int{2, 3, 1, 1, 4})
	if test1 != 2 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := jump([]int{2, 3, 0, 1, 4})
	if test2 != 2 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
