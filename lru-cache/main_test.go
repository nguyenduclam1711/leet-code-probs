package lrucache

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lru := Constructor(3)

	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)

	value1 := lru.Get(4)
	if value1 != 4 {
		fmt.Println("value1", value1)
		t.FailNow()
	}

	value2 := lru.Get(3)
	if value2 != 3 {
		fmt.Println("value2", value2)
		t.FailNow()
	}

	value3 := lru.Get(2)
	if value3 != 2 {
		fmt.Println("value1", value3)
		t.FailNow()
	}

	value4 := lru.Get(1)
	if value4 != -1 {
		fmt.Println("value4", value4)
		t.FailNow()
	}

	lru.Put(5, 5)

	value5 := lru.Get(1)
	if value5 != -1 {
		fmt.Println("value5", value5)
		t.FailNow()
	}

	value6 := lru.Get(2)
	if value6 != 2 {
		fmt.Println("value6", value6)
		t.FailNow()
	}

	value7 := lru.Get(3)
	if value7 != 3 {
		fmt.Println("value7", value7)
		t.FailNow()
	}

	value8 := lru.Get(4)
	if value8 != -1 {
		fmt.Println("value8", value8)
		t.FailNow()
	}

	value9 := lru.Get(5)
	if value9 != 5 {
		fmt.Println("value9", value9)
		t.FailNow()
	}
}
