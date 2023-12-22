package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	depthL := dfs(node.Left, depth+1)
	depthR := dfs(node.Right, depth+1)

	if depthL > depthR {
		return depthL
	}
	return depthR
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MaxDepth(root *TreeNode) int {
	return dfs(root, 0)
}
