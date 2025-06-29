package canplaceflowers

import (
	"fmt"
	"testing"
)

func TestCanPlaceFlowers(t *testing.T) {
	test1 := canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)

	if !test1 {
		fmt.Println("test1", test1)
		t.Fail()
	}

	test2 := canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2)

	if test2 {
		fmt.Println("test2", test2)
		t.Fail()
	}

	test3 := canPlaceFlowers([]int{0, 0, 1, 0, 0}, 1)

	if !test3 {
		fmt.Println("test3", test3)
		t.Fail()
	}

	test4 := canPlaceFlowers([]int{0, 0, 0, 0, 0, 1, 0, 0}, 0)

	if !test4 {
		fmt.Println("test4", test4)
		t.Fail()
	}
}
