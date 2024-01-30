package maxconsecutiveonesiii

import (
	"fmt"
	"testing"
)

func TestLongestOnes(t *testing.T) {
	test1 := longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)
	if test1 != 6 {
		fmt.Println("test1", test1)
		t.Fail()
	}

	test2 := longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
	if test2 != 10 {
		fmt.Println("test2", test2)
		t.Fail()
	}
}
