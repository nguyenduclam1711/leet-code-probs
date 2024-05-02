package textjustification

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFullJustify(t *testing.T) {
	test1 := fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
	if !reflect.DeepEqual([]string{
		"This    is    an",
		"example  of text",
		"justification.  ",
	}, test1) {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16)
	if !reflect.DeepEqual([]string{
		"What   must   be",
		"acknowledgment  ",
		"shall be        ",
	}, test2) {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20)
	if !reflect.DeepEqual([]string{
		"Science  is  what we",
		"understand      well",
		"enough to explain to",
		"a  computer.  Art is",
		"everything  else  we",
		"do                  ",
	}, test3) {
		fmt.Println("test3", test3)
		t.FailNow()
	}
}
