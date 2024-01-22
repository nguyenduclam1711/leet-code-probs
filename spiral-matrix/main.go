package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	topWall := 0
	rightWall := len(matrix[0]) - 1
	bottomWall := len(matrix)
	leftWall := -1

	res := []int{}
	x, y := 0, 0
	numberOfEle := len(matrix) * len(matrix[0])

	for len(res) < numberOfEle {
		res = append(res, matrix[x][y])

		if x == topWall && y < rightWall {
			// go right
			y++
			if y == rightWall {
				bottomWall--
			}
			continue
		}
		if y == rightWall && x < bottomWall {
			// go down
			x++
			if x == bottomWall {
				leftWall++
			}
			continue
		}
		if x == bottomWall && y > leftWall {
			// go left
			y--
			if y == leftWall {
				topWall++
			}
			continue
		}
		if y == leftWall && x > topWall {
			// go up
			x--
			if x == topWall {
				rightWall--
			}
			continue
		}
	}
	return res
}
