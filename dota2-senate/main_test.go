package dota2senate

import (
	"fmt"
	"testing"
)

func TestPredictPartyVictory(t *testing.T) {
	test1 := predictPartyVictory("RD")
	if test1 != "Radiant" {
		fmt.Println("test1", test1)
		t.FailNow()
	}
	test2 := predictPartyVictory("RDD")
	if test2 != "Dire" {
		fmt.Println("test2", test2)
		t.FailNow()
	}
	test3 := predictPartyVictory("DRRDRDRDRDDRDRDR")
	if test3 != "Radiant" {
		fmt.Println("test3", test3)
		t.FailNow()
	}
	test4 := predictPartyVictory("RDRDRDD")
	if test4 != "Radiant" {
		fmt.Println("test4", test4)
		t.FailNow()
	}
}
