var getStackMapKey = function(p) {
  return `${p.x},${p.y}`
}

var swapMatrix = function(matrix, arr, moveTimes) {
	const stack = {}

  for (let i = 0; i < arr.length; i++) {
    const pointPos = arr[i]
		const swapPos = (i + arr.length - moveTimes) % arr.length
		const stackKey = getStackMapKey(pointPos)
		stack[stackKey] = matrix[pointPos.x][pointPos.y]
		const swapPosPoint = arr[swapPos]
		if (i < moveTimes) {
			matrix[pointPos.x][pointPos.y] = matrix[swapPosPoint.x][swapPosPoint.y]
		} else {
			matrix[pointPos.x][pointPos.y] = stack[getStackMapKey(swapPosPoint)]
		}
	}
}

var rotate = function(matrix) {
	let matrixLen = matrix.length
	let loop = 0

	while (matrixLen > 0) {
		// totalNumberOfOuterLinePos = perimeter - 4 (4 is the number of corners)
		// totalNumberOfOuterLinePos := ((matrixLen + matrixLen) * 2) - 4
		const totalNumberOfOuterLinePos = matrixLen*4 - 4
		if (totalNumberOfOuterLinePos > 0) {
      const outerMatrixLines = Array(totalNumberOfOuterLinePos)
      let x = 0
      let y = 0
      for (let i = 0; i < totalNumberOfOuterLinePos; i++) {
        outerMatrixLines[i] = {
          x: x + loop,
          y: y + loop,
        }
				if (x <= y) {
					if (x === matrixLen-1 && y === matrixLen-1) {
						y--
						continue
					}
					if (y < matrixLen-1) {
						y++
					} else {
						x++
					}
				} else {
					if (y > 0) {
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

const matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
rotate(matrix)
console.log(matrix)
