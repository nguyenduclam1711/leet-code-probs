package guessnumberhigherorlower

import "math/rand"

var randomNum = rand.Int()

func guess(n int) int {
	if n == randomNum {
		return 0
	}
	if n > randomNum {
		return -1
	}
	return 1
}

func GuessNumber(n int) int {
	l, r := 1, n

	for l <= r {
		m := (l + r) / 2
		guessRes := guess(m)
		if guessRes == 0 {
			return m
		} else if guessRes > 0 {
			l = m + 1
		} else {
			r = m
		}
	}
	return -1
}
