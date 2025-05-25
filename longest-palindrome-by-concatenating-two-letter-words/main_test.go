package longestpalindromebyconcatenatingtwoletterwords

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	test1 := longestPalindrome([]string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"})
	if test1 != 22 {
		fmt.Println("test1", test1)
		t.Fail()
	}
	test2 := longestPalindrome([]string{"em", "pe", "mp", "ee", "pp", "me", "ep", "em", "em", "me"})
	if test2 != 14 {
		fmt.Println("test2", test2)
		t.Fail()
	}
}
