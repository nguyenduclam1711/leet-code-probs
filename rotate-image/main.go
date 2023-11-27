package main

import "fmt"

func getSwapPosition(x int, y int, matrixLen int) (int, int) {
	maxPos := matrixLen - 1
	// top border
	if x == 0 {
		return maxPos - y, 0
	}
	// right border
	if y == maxPos {
		return 0, x
	}
	// bottom border
	if x == maxPos {
		return maxPos - y, maxPos
	}
	// left border
	if y == 0 {
		return maxPos, x
	}
	return -1, -1
}

func getMapKey(x int, y int) string {
	return fmt.Sprint(x, ",", y)
}

func rotate(matrix [][]int) {
	matrixLen := len(matrix)
	loop := 0
	m := map[string]int{}

	for matrixLen > 0 {
		totalNumberOfOuterLinePos := matrixLen*4 - 4
		if totalNumberOfOuterLinePos > 0 {
			for i, x, y := 0, 0, 0; i < totalNumberOfOuterLinePos; i++ {
				swapX, swapY := getSwapPosition(x, y, matrixLen)
				posX, posY := x+loop, y+loop
				posSwapX, posSwapY := swapX+loop, swapY+loop
				m[getMapKey(posX, posY)] = matrix[posX][posY]

				if _, found := m[getMapKey(posSwapX, posSwapY)]; found {
					matrix[posX][posY] = m[getMapKey(posSwapX, posSwapY)]
				} else {
					matrix[posX][posY] = matrix[posSwapX][posSwapY]
				}
				if x <= y {
					if x == y && y == matrixLen-1 {
						y--
						continue
					}
					if y < matrixLen-1 {
						y++
					} else {
						x++
					}
				} else {
					if y > 0 {
						y--
					} else {
						x--
					}
				}
			}
		}
		// the inside outer line always smaller than the larger on by 2
		matrixLen -= 2
		loop++
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	// matrix := [][]int{
	// 	{5, 1, 9, 11},
	// 	{2, 4, 8, 10},
	// 	{13, 3, 6, 7},
	// 	{15, 14, 12, 16},
	// }
	rotate(matrix)
	fmt.Println(matrix)
}
