package subtreeofanothertree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSame(node *TreeNode, subNode *TreeNode) bool {
	if node == nil && subNode == nil {
		return true
	}
	if node == nil || subNode == nil {
		return false
	}
	if node.Val != subNode.Val {
		return false
	}
	return checkSame(node.Left, subNode.Left) && checkSame(node.Right, subNode.Right)
}

func dfs(node *TreeNode, subRoot *TreeNode) bool {
	if node == nil {
		return false
	}
	if node.Val == subRoot.Val {
		isSame := checkSame(node, subRoot)
		if isSame {
			return isSame
		}
	}
	return dfs(node.Left, subRoot) || dfs(node.Right, subRoot)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return dfs(root, subRoot)
}
