package kthsmallestelementinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorder traverse
func dfs(node *TreeNode, s *[]int, k int) {
	if node == nil {
		return
	}
	dfs(node.Left, s, k)
	*s = append(*s, node.Val)
	if len(*s) == k {
		return
	}
	dfs(node.Right, s, k)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func KthSmallest(root *TreeNode, k int) int {
	s := []int{}
	dfs(root, &s, k)
	return s[k-1]
}
