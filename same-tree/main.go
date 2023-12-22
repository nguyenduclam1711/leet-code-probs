package sametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recurse(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return recurse(p.Left, q.Left) && recurse(p.Right, q.Right)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return recurse(p, q)
}
