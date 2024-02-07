package searchinabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		return SearchBST(root.Right, val)
	}
	return SearchBST(root.Left, val)
}
