package lowestcommonancestorofabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func walk(node *TreeNode, min int, max int) *TreeNode {
	if node.Val > max {
		return walk(node.Left, min, max)
	} else if node.Val < min {
		return walk(node.Right, min, max)
	} else {
		return node
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	max := p.Val
	min := q.Val
	if q.Val > p.Val {
		max = q.Val
		min = p.Val
	}
	return walk(root, min, max)
}
