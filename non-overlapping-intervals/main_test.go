package nonoverlappingintervals

import (
	"fmt"
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	test1 := eraseOverlapIntervals([][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	})

	if test1 != 1 {
		fmt.Println("test1", test1)
		t.FailNow()
	}

	test2 := eraseOverlapIntervals([][]int{
		{-52, 31},
		{-73, -26},
		{82, 97},
		{-65, -11},
		{-62, -49},
		{95, 99},
		{58, 95},
		{-31, 49},
		{66, 98},
		{-63, 2},
		{30, 47},
		{-40, -26},
	})

	if test2 != 7 {
		fmt.Println("test2", test2)
		t.FailNow()
	}

	test3 := eraseOverlapIntervals([][]int{
		{-3035, 30075},
		{1937, 6906},
		{11834, 20971},
		{44578, 45600},
		{28565, 37578},
		{19621, 34415},
		{32985, 36313},
		{-8144, 1080},
		{-15279, 21851},
		{-27140, -14703},
		{-12098, 16264},
		{-36057, -16287},
		{47939, 48626},
		{-15129, -5773},
		{10508, 46685},
		{-35323, -26257},
	})

	if test3 != 9 {
		t.FailNow()
		fmt.Println("test3", test3)
	}
}
