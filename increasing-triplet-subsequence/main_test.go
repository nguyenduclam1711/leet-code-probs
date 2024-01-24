package increasingtripletsubsequence

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	test1 := increasingTriplet([]int{1, 2, 3, 4, 5})
	if !test1 {
		t.Fail()
	}

	test2 := increasingTriplet([]int{5, 4, 3, 2, 1})
	if test2 {
		t.Fail()
	}

	test3 := increasingTriplet([]int{2, 1, 5, 0, 4, 6})
	if !test3 {
		t.Fail()
	}
}
