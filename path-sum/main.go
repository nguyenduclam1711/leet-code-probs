package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recurse(node *TreeNode, targetSum int, prevSum int) bool {
	if node == nil {
		return false
	}
	currSum := prevSum + node.Val

	if node.Left == nil && node.Right == nil && currSum == targetSum {
		return true
	}
	return recurse(node.Left, targetSum, currSum) || recurse(node.Right, targetSum, currSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return recurse(root, targetSum, 0)
}
