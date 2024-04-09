package dailytemperatures

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	test1 := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	if !reflect.DeepEqual(test1, []int{1, 1, 4, 2, 1, 1, 0, 0}) {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := dailyTemperatures([]int{30, 40, 50, 60})
	if !reflect.DeepEqual(test2, []int{1, 1, 1, 0}) {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := dailyTemperatures([]int{30, 60, 90})
	if !reflect.DeepEqual(test3, []int{1, 1, 0}) {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
