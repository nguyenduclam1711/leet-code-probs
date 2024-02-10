package nearestexitfromentranceinmaze

import (
	"math"
)

func nearestExit(maze [][]byte, entrance []int) int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	res := math.MaxInt
	q := [][]int{
		entrance,
	}

	addToQueue := func(x int, y int, prevStep int) {
		if x < 0 || x == len(maze) || y < 0 || y == len(maze[0]) {
			return
		}
		if steps[x][y] > 0 || (x == entrance[0] && y == entrance[1]) {
			return
		}
		if maze[x][y] == '+' {
			return
		}
		currStep := prevStep + 1
		steps[x][y] = currStep
		if x == 0 || x == len(maze)-1 || y == 0 || y == len(maze[0])-1 {
			if currStep < res {
				res = currStep
			}
			return
		}
		q = append(q, []int{x, y})
	}

	for len(q) > 0 {
		currQueueLen := len(q)
		for _, pos := range q {
			x, y := pos[0], pos[1]
			// go up
			addToQueue(x-1, y, steps[x][y])
			// go right
			addToQueue(x, y+1, steps[x][y])
			// go down
			addToQueue(x+1, y, steps[x][y])
			// go left
			addToQueue(x, y-1, steps[x][y])
		}
		q = q[currQueueLen:]
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
