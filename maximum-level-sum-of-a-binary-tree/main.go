package maximumlevelsumofabinarytree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	q := []*TreeNode{root}
	maxSum := math.MinInt
	level := 0
	res := 0

	for len(q) > 0 {
		sum := 0
		level++
		newQ := []*TreeNode{}
		for _, n := range q {
			sum += n.Val
			if n.Left != nil {
				newQ = append(newQ, n.Left)
			}
			if n.Right != nil {
				newQ = append(newQ, n.Right)
			}
		}
		if sum > maxSum {
			maxSum = sum
			res = level
		}
		q = newQ
	}
	return res
}
