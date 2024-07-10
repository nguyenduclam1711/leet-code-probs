package countcompletetreenodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recurse(node *TreeNode, count *int) {
	if node == nil {
		return
	}
	*count++
	recurse(node.Left, count)
	recurse(node.Right, count)
}

func countNodes(root *TreeNode) int {
	count := 0
	recurse(root, &count)
	return count
}
