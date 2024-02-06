package countgoodnodesinbinarytree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, nums *[]int, res *int) {
	if node == nil {
		return
	}
	isAppendToNums := false
	if node.Val >= (*nums)[len(*nums)-1] {
		isAppendToNums = true
		*res++
		*nums = append(*nums, node.Val)
	}
	dfs(node.Left, nums, res)
	dfs(node.Right, nums, res)
	if isAppendToNums {
		*nums = (*nums)[:len(*nums)-1]
	}
}

func GoodNodes(root *TreeNode) int {
	maxNums := []int{math.MinInt}
	res := 0
	dfs(root, &maxNums, &res)
	return res
}
