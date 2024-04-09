package onlinestockspan

import (
	"fmt"
	"testing"
)

func TestOnlineStockSpan(t *testing.T) {
	stockSpanner := Constructor()
	test1 := stockSpanner.Next(100)
	if test1 != 1 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := stockSpanner.Next(80)
	if test2 != 1 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := stockSpanner.Next(60)
	if test3 != 1 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := stockSpanner.Next(70)
	if test4 != 2 {
		fmt.Println("test4", test4)
		t.FailNow()
	}
	test5 := stockSpanner.Next(60)
	if test5 != 1 {
		fmt.Println("test5", test5)
		t.FailNow()
	}
	test6 := stockSpanner.Next(75)
	if test6 != 4 {
		fmt.Println("test6", test6)
		t.FailNow()
	}
	test7 := stockSpanner.Next(85)
	if test7 != 6 {
		fmt.Println("test7", test7)
		t.FailNow()
	}
}
