package decodestring

import (
	"fmt"
	"testing"
)

func TestDecodeString(t *testing.T) {
	test1 := decodeString("3[a]2[bc]")
	if test1 != "aaabcbc" {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := decodeString("3[a2[c]]")
	if test2 != "accaccacc" {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := decodeString("2[abc]3[cd]ef")
	if test3 != "abcabccdcdcdef" {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := decodeString("2[a2[b]2[c2[d]]]")
	if test4 != "abbcddcddabbcddcdd" {
		fmt.Println("test4", test4)
		t.FailNow()
	}
	test5 := decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef")
	if test5 != "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef" {
		fmt.Println("test5", test5)
		t.FailNow()
	}
	test6 := decodeString("pp2[c]")
	if test6 != "ppcc" {
		fmt.Println("test6", test6)
		t.FailNow()
	}
}
