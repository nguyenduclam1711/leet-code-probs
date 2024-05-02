package substringwithconcatenationofallwords

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	// test1 := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	// if !reflect.DeepEqual([]int{0, 9}, test1) {
	// 	fmt.Println("test1", test1)
	// 	t.FailNow()
	// }
	// test2 := findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})
	// if !reflect.DeepEqual([]int{}, test2) {
	// 	fmt.Println("test2", test2)
	// 	t.FailNow()
	// }
	// test3 := findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})
	// if !reflect.DeepEqual([]int{6, 9, 12}, test3) {
	// 	fmt.Println("test3", test3)
	// 	t.FailNow()
	// }
	test4 := findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})
	if !reflect.DeepEqual([]int{8}, test4) {
		fmt.Println("test4", test4)
		t.FailNow()
	}
}
