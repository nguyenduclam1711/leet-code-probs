package smallestnumberininfiniteset

import (
	"testing"
)

func TestSmallestNumberInInfiniteSet(t *testing.T) {
	h := Constructor()

	h.AddBack(2)
	v1 := h.PopSmallest()
	if v1 != 1 {
		t.Fatal("v1", v1)
	}
	v2 := h.PopSmallest()
	if v2 != 2 {
		t.Fatal("v2", v2)
	}
	v3 := h.PopSmallest()
	if v3 != 3 {
		t.Fatal("v3", v3)
	}
	h.AddBack(1)
	v4 := h.PopSmallest()
	if v4 != 1 {
		t.Fatal("v4", v4)
	}
	v5 := h.PopSmallest()
	if v5 != 4 {
		t.Fatal("v5", v5)
	}
	v6 := h.PopSmallest()
	if v6 != 5 {
		t.Fatal("v6", v6)
	}
}
