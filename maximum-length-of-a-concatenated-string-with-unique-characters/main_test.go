package maximumlengthofaconcatenatedstringwithuniquecharacters

import (
	"fmt"
	"testing"
)

func TestMaxLength(t *testing.T) {
	test1 := maxLength([]string{"un", "iq", "ue"})
	if test1 != 4 {
		fmt.Println("test1", test1)
		t.Fail()
	}

	test2 := maxLength([]string{"cha", "r", "act", "ers"})
	if test2 != 6 {
		fmt.Println("test2", test2)
		t.Fail()
	}

	test3 := maxLength([]string{"abcdefghijklmnopqrstuvwxyz"})
	if test3 != 26 {
		fmt.Println("test3", test3)
		t.Fail()
	}

	test4 := maxLength([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"})
	if test4 != 16 {
		fmt.Println("test4", test4)
		t.Fail()
	}

	test5 := maxLength([]string{"abcdefghijklm", "bcdefghijklmn", "cdefghijklmno", "defghijklmnop", "efghijklmnopq", "fghijklmnopqr", "ghijklmnopqrs", "hijklmnopqrst", "ijklmnopqrstu", "jklmnopqrstuv", "klmnopqrstuvw", "lmnopqrstuvwx", "mnopqrstuvwxy", "nopqrstuvwxyz", "opqrstuvwxyza", "pqrstuvwxyzab"})
	if test5 != 26 {
		fmt.Println("test5", test5)
		t.Fail()
	}
}
