package main

import "fmt"

type point struct {
	x int
	y int
}

func getStackMapKey(p point) string {
	return fmt.Sprint(p.x, ",", p.y)
}

func swapMatrix(matrix [][]int, arr []point, moveTimes int) {
	stack := map[string]int{}

	for i, pointPos := range arr {
		swapPos := (i + len(arr) - moveTimes) % len(arr)
		stackKey := getStackMapKey(pointPos)
		stack[stackKey] = matrix[pointPos.x][pointPos.y]
		swapPosPoint := arr[swapPos]
		if i < moveTimes {
			matrix[pointPos.x][pointPos.y] = matrix[swapPosPoint.x][swapPosPoint.y]
		} else {
			matrix[pointPos.x][pointPos.y] = stack[getStackMapKey(swapPosPoint)]
		}
	}
}

func rotate(matrix [][]int) {
	matrixLen := len(matrix)
	loop := 0

	for matrixLen > 0 {
		// totalNumberOfOuterLinePos = perimeter - 4 (4 is the number of corners)
		// totalNumberOfOuterLinePos := ((matrixLen + matrixLen) * 2) - 4
		totalNumberOfOuterLinePos := matrixLen*4 - 4
		if totalNumberOfOuterLinePos > 0 {
			outerMatrixLines := []point{}
			for i, x, y := 0, 0, 0; i < totalNumberOfOuterLinePos; i++ {
				outerMatrixLines = append(outerMatrixLines, point{x + loop, y + loop})
				if x <= y {
					if x == matrixLen-1 && y == matrixLen-1 {
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
			swapMatrix(matrix, outerMatrixLines, matrixLen-1)
		}
		// the inside outer line always smaller than the larger on by 2
		matrixLen -= 2
		loop++
	}
}

func main() {
	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
