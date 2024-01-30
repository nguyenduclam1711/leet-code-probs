package maximumnumberofvowelsinasubstringofgivenlength

import (
	"fmt"
	"testing"
)

func TestMaxVowels(t *testing.T) {
	test1 := maxVowels("abciiidef", 3)
	if test1 != 3 {
		fmt.Println("test1", test1)
		t.Fail()
	}

	test2 := maxVowels("aeiou", 2)
	if test2 != 2 {
		fmt.Println("test2", test2)
		t.Fail()
	}

	test3 := maxVowels("leetcode", 3)
	if test3 != 2 {
		fmt.Println("test3", test3)
		t.Fail()
	}
}
