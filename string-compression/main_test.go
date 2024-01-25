package stringcompression

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCompress(t *testing.T) {
	s1 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	test1 := compress(s1)

	if !reflect.DeepEqual(s1[:test1], []byte{'a', '2', 'b', '2', 'c', '3'}) {
		fmt.Println("test1", test1, string(s1[:test1]))
		t.Fail()
	}

	s2 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	test2 := compress(s2)

	if !reflect.DeepEqual(s2[:test2], []byte{'a', 'b', '1', '2'}) {
		fmt.Println("test2", test2, string(s2[:test2]))
		t.Fail()
	}

	s3 := []byte{'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o'}
	test3 := compress(s3)

	if !reflect.DeepEqual(s3[:test3], []byte{'o', '1', '0'}) {
		fmt.Println("test3", test3, string(s3[:test3]))
		t.Fail()
	}
}
