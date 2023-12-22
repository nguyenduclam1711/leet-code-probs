package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recurse(node *TreeNode) {
	if node == nil {
		return
	}
	node.Left, node.Right = node.Right, node.Left
	recurse(node.Left)
	recurse(node.Right)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func InvertTree(root *TreeNode) *TreeNode {
	recurse(root)
	return root
}
