package insertdeletegetrandomo1

import (
	"fmt"
	"testing"
)

func test1(t *testing.T) {
	rs := Constructor()
	test1 := rs.Insert(1)
	if !test1 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := rs.Remove(2)
	if test2 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := rs.Insert(2)
	if !test3 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := rs.GetRandom()
	if test4 != 1 && test4 != 2 {
		fmt.Println("test4", test4)
		t.FailNow()
	}
	test5 := rs.Remove(1)
	if !test5 {
		fmt.Println("test5", test5)
		t.FailNow()
	}
	test6 := rs.Insert(2)
	if test6 {
		fmt.Println("test6", test6)
		t.FailNow()
	}
	test7 := rs.GetRandom()
	if test7 != 2 {
		fmt.Println("test7", test7)
		t.FailNow()
	}
}

func test2(t *testing.T) {
	rs := Constructor()
	test1 := rs.Insert(0)
	if !test1 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := rs.Insert(1)
	if !test2 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := rs.Remove(0)
	if !test3 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := rs.Insert(2)
	if !test4 {
		fmt.Println("test4", test4)
		t.FailNow()
	}
	test5 := rs.Remove(1)
	if !test5 {
		fmt.Println("test5", test5)
		t.FailNow()
	}
	test6 := rs.GetRandom()
	if test6 != 2 {
		fmt.Println("test6", test6)
		t.FailNow()
	}
}

func TestRandomizedSet(t *testing.T) {
	test1(t)
	test2(t)
}
