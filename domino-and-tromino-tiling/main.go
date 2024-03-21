package dominoandtrominotiling

import "math"

// Couldn't find a solution so this is the solution that I find understandable
// https://leetcode.com/problems/domino-and-tromino-tiling/solutions/4723561/simple-dynamic-programming-solution-with-graph-analysis/?envType=study-plan-v2&envId=leetcode-75
func numTilings(n int) int {
	mod := int(math.Pow(10, 9)) + 7
	filled, upperEmpty, lowerEmpty, empty := 1, 0, 0, 0
	for i := 1; i <= n; i++ {
		// When the last collumn is filled
		newFilled := (filled + upperEmpty + lowerEmpty + empty) % mod
		// When the upper cell of the last collumn is empty
		newUpperEmpty := (lowerEmpty + empty) % mod
		// When the lower cell of the last collumn is empty
		newLowerEmpty := (upperEmpty + empty) % mod
		// When the last cell is empty
		newEmpty := filled % mod

		filled, upperEmpty, lowerEmpty, empty = newFilled, newUpperEmpty, newLowerEmpty, newEmpty
	}
	return filled
}
