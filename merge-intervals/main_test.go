package mergeintervals

import (
	"testing"
)

func TestMerge(t *testing.T) {
	res := merge([][]int{
		{1, 4},
		{0, 5},
	})

	if res[0][0] != 0 || res[0][1] != 5 {
		t.Fail()
	}
}
