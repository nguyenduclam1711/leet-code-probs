package determineiftwostringsareclose

import "testing"

func TestCloseStrings(t *testing.T) {
	test1 := closeStrings("abc", "bca")
	if !test1 {
		t.Fail()
	}
}
