package summaryranges

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	test1 := summaryRanges([]int{0, 1, 2, 4, 5, 7})
	if !reflect.DeepEqual(test1, []string{"0->2", "4->5", "7"}) {
		fmt.Println("test1", test1)
		t.FailNow()
	}

	test2 := summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})
	if !reflect.DeepEqual(test2, []string{"0", "2->4", "6", "8->9"}) {
		fmt.Println("test2", test2)
		t.FailNow()
	}
}
