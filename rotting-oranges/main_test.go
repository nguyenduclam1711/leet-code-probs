package rottingoranges

import (
	"fmt"
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	// test1 := orangesRotting([][]int{
	// 	{2, 1, 1},
	// 	{1, 1, 0},
	// 	{0, 1, 1},
	// })
	// if test1 != 4 {
	// 	fmt.Println("test1", test1)
	// 	t.FailNow()
	// }
	// test2 := orangesRotting([][]int{
	// 	{2, 1, 1},
	// 	{0, 1, 1},
	// 	{1, 0, 1},
	// })
	// if test2 != -1 {
	// 	fmt.Println("test2", test2)
	// 	t.FailNow()
	// }
	// test3 := orangesRotting([][]int{
	// 	{0, 2},
	// })
	// if test3 != 0 {
	// 	fmt.Println("test3", test3)
	// 	t.FailNow()
	// }
	test4 := orangesRotting([][]int{
		{0},
	})
	if test4 != 0 {
		fmt.Println("test4", test4)
		t.FailNow()
	}
}
