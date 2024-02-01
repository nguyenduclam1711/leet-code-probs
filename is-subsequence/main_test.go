package issubsequence

import "testing"

func TestIsSubsequence(t *testing.T) {
	test1 := isSubsequence("abc", "ahbgdc")
	if !test1 {
		t.FailNow()
	}
	test2 := isSubsequence("axc", "ahbgdc")
	if test2 {
		t.FailNow()
	}
}
