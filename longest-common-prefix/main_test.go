package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	test1 := longestCommonPrefix([]string{"a"})
	if test1 != "a" {
		t.FailNow()
	}
	test2 := longestCommonPrefix([]string{"flower", "flow", "flight"})
	if test2 != "fl" {
		t.FailNow()
	}
	test3 := longestCommonPrefix([]string{"dog", "racecar", "car"})
	if test3 != "" {
		t.FailNow()
	}
}
