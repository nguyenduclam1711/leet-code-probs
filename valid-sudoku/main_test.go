package validsudoku

import (
	"fmt"
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	board1String := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	board1 := [][]byte{}
	for _, strs := range board1String {
		board1 = append(board1, []byte{})
		for _, str := range strs {
			board1[len(board1)-1] = append(board1[len(board1)-1], str[0])
		}
	}
	test1 := isValidSudoku(board1)
	if !test1 {
		fmt.Println("test1", test1)
		t.FailNow()
	}

	board2String := [][]string{
		{"8", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	board2 := [][]byte{}
	for _, strs := range board2String {
		board2 = append(board2, []byte{})
		for _, str := range strs {
			board2[len(board2)-1] = append(board2[len(board2)-1], str[0])
		}
	}
	test2 := isValidSudoku(board2)
	if test2 {
		fmt.Println("test2", test2)
		t.FailNow()
	}

	board3String := [][]string{
		{".", ".", ".", ".", "5", ".", ".", "1", "."},
		{".", "4", ".", "3", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "3", ".", ".", "1"},
		{"8", ".", ".", ".", ".", ".", ".", "2", "."},
		{".", ".", "2", ".", "7", ".", ".", ".", "."},
		{".", "1", "5", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "2", ".", ".", "."},
		{".", "2", ".", "9", ".", ".", ".", ".", "."},
		{".", ".", "4", ".", ".", ".", ".", ".", "."},
	}
	board3 := [][]byte{}
	for _, strs := range board3String {
		board3 = append(board3, []byte{})
		for _, str := range strs {
			board3[len(board3)-1] = append(board3[len(board3)-1], str[0])
		}
	}
	test3 := isValidSudoku(board3)
	if test3 {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
