package editdistance

import (
	"fmt"
	"testing"
)

func TestMinDistance(t *testing.T) {
	test1 := minDistance("horse", "ros")
	if test1 != 3 {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := minDistance("intention", "execution")
	if test2 != 5 {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
