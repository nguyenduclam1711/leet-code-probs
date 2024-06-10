package longestconsecutivesequence

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	test1 := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	if test1 != 4 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	if test2 != 9 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
