package longestsubarrayof1safterdeletingoneelement

import (
	"fmt"
	"testing"
)

func TestLongestSubarray(t *testing.T) {
	test1 := longestSubarray([]int{1, 1, 0, 1})
	if test1 != 3 {
		fmt.Println("test1", test1)
		t.FailNow()
	}

	test2 := longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1})
	if test2 != 5 {
		fmt.Println("test2", test2)
		t.FailNow()
	}

	test3 := longestSubarray([]int{1, 1, 1})
	if test3 != 2 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
