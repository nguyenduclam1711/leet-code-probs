package courseschedule

import (
	"testing"
)

func TestCanFinish(t *testing.T) {
	if !canFinish(2, [][]int{{1, 0}}) {
		t.Fail()
	}
	if canFinish(2, [][]int{{1, 0}, {0, 1}}) {
		t.Fail()
	}
	if !canFinish(5, [][]int{
		{1, 4},
		{2, 4},
		{3, 1},
		{3, 2},
	}) {
		t.Fail()
	}
	if !canFinish(1, [][]int{}) {
		t.Fail()
	}
	if canFinish(20, [][]int{
		{0, 10},
		{3, 18},
		{5, 5},
		{6, 11},
		{11, 14},
		{13, 1},
		{15, 1},
		{17, 4},
	}) {
		t.Fail()
	}
	if canFinish(3, [][]int{
		{0, 2},
		{1, 2},
		{2, 0},
	}) {
		t.Fail()
	}
	if canFinish(8, [][]int{
		{1, 0},
		{2, 6},
		{1, 7},
		{5, 1},
		{6, 4},
		{7, 0},
		{0, 5},
	}) {
		t.Fail()
	}
}
