package interleavingstring

import (
	"fmt"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	test1 := isInterleave("aabd", "abdc", "aabdabcd")
	if !test1 {
		fmt.Println("test1", test1)
		t.Fail()
	}
	test2 := isInterleave("abbbbbbcabbacaacccababaabcccabcacbcaabbbacccaaaaaababbbacbb", "ccaacabbacaccacababbbbabbcacccacccccaabaababacbbacabbbbabc", "cacbabbacbbbabcbaacbbaccacaacaacccabababbbababcccbabcabbaccabcccacccaabbcbcaccccaaaaabaaaaababbbbacbbabacbbacabbbbabc")
	if !test2 {
		fmt.Println("test2", test2)
		t.Fail()
	}
}
