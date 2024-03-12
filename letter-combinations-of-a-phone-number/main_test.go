package lettercombinationsofaphonenumber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	test1 := letterCombinations("23")
	if !reflect.DeepEqual(test1, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}) {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := letterCombinations("")
	if !reflect.DeepEqual(test2, []string{}) {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := letterCombinations("2")
	if !reflect.DeepEqual(test3, []string{"a", "b", "c"}) {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
