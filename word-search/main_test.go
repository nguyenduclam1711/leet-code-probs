package wordsearch

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	test1 := exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED")

	if !test1 {
		fmt.Println("test1", test1)
		t.Fail()
	}

	test2 := exist([][]byte{
		{'a'},
	}, "a")

	if !test2 {
		fmt.Println("test2", test2)
		t.Fail()
	}

	test3 := exist([][]byte{
		{'a', 'b'},
	}, "ba")

	if !test3 {
		fmt.Println("test3", test3)
		t.Fail()
	}
}
