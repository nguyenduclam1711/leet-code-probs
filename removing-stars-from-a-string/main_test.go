package removingstarsfromastring

import (
	"fmt"
	"testing"
)

func TestRemoveStars(t *testing.T) {
	test1 := removeStars("leet**cod*e")
	if test1 != "lecoe" {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := removeStars("erase*****")
	if test2 != "" {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
