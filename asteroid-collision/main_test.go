package asteroidcollision

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	test1 := asteroidCollision([]int{5, 10, -5})
	if !reflect.DeepEqual(test1, []int{5, 10}) {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := asteroidCollision([]int{8, -8})
	if !reflect.DeepEqual(test2, []int{}) {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := asteroidCollision([]int{10, 2, -5})
	if !reflect.DeepEqual(test3, []int{10}) {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := asteroidCollision([]int{-2, -1, 1, 2})
	if !reflect.DeepEqual(test4, []int{-2, -1, 1, 2}) {
		fmt.Println("test4", test4)
		t.FailNow()
	}
}
