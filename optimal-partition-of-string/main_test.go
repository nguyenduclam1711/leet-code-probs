package optimalpartitionofstring

import (
	"fmt"
	"testing"
)

func TestPartitionString(t *testing.T) {
	test1 := partitionString("abacaba")
	if test1 != 4 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := partitionString("ssssss")
	if test2 != 6 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
