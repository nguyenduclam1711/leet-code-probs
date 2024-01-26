package findthehighestaltitude

import "testing"

func TestLargestAltitude(t *testing.T) {
	test1 := largestAltitude([]int{-5, 1, 5, 0, -7})

	if test1 != 1 {
		t.Fail()
	}
}
