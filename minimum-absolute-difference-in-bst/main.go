package minimumabsolutedifferenceinbst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMostLeftNode(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}
	if node.Left != nil {
		return getMostLeftNode(node.Left)
	}
	return node
}

func getMostRightNode(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}
	if node.Right != nil {
		return getMostRightNode(node.Right)
	}
	return node
}

func dfs(node *TreeNode, res *int) {
	if node == nil {
		return
	}
	closestRight := getMostLeftNode(node.Right)
	closestLeft := getMostRightNode(node.Left)

	if closestRight != nil {
		newRes := closestRight.Val - node.Val
		if newRes < *res {
			*res = newRes
		}
	}
	if closestLeft != nil {
		newRes := node.Val - closestLeft.Val
		if newRes < *res {
			*res = newRes
		}
	}
	dfs(node.Left, res)
	dfs(node.Right, res)
}

func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt
	dfs(root, &res)
	return res
}
