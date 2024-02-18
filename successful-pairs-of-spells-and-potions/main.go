package successfulpairsofspellsandpotions

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := make([]int, len(spells))

	for i, spell := range spells {
		l, r := 0, len(potions)
		foundTotal := false
		for l < r {
			if potions[l]*spell >= int(success) {
				res[i] = len(potions) - l
				foundTotal = true
				break
			}
			m := (l + r) / 2
			power := potions[m] * spell
			if power == int(success) {
				res[i] = len(potions) - m
				break
			} else if power > int(success) {
				res[i] = len(potions) - m
				r = m
			} else {
				l = m + 1
			}
		}
		if !foundTotal {
			m := (l + r) / 2
			for j := m - 1; j >= 0; j-- {
				if spell*potions[j] < int(success) {
					break
				}
				res[i]++
			}
		}
	}
	return res
}
