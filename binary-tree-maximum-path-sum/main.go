package binarytreemaximumpathsum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	var recurse func(*TreeNode) int

	recurse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			if node.Val > res {
				res = node.Val
			}
			return node.Val
		}
		maxLeft := recurse(node.Left)
		maxRight := recurse(node.Right)

		totalPathRight := node.Val + maxRight
		totalPathLeft := node.Val + maxLeft
		currNodeMax := max(node.Val, totalPathLeft, totalPathRight)

		if currNodeMax == node.Val {
			if currNodeMax > res {
				res = currNodeMax
			}
		} else if currNodeMax == totalPathLeft {
			newRes := max(currNodeMax+maxRight, currNodeMax)
			if newRes > res {
				res = newRes
			}
		} else if currNodeMax == totalPathRight {
			newRes := max(currNodeMax+maxLeft, currNodeMax)
			if newRes > res {
				res = newRes
			}
		}
		return currNodeMax
	}

	recurse(root)
	return res
}
