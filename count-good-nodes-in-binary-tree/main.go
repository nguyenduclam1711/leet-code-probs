package countgoodnodesinbinarytree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, nums *[]int, res *int) bool {
	if node == nil {
		return false
	}
	isGreater := false
	if node.Val >= (*nums)[len(*nums)-1] {
		*res++
		*nums = append(*nums, node.Val)
		isGreater = true
	}
	isGreaterLeft := dfs(node.Left, nums, res)
	if isGreaterLeft {
		*nums = (*nums)[:len(*nums)-1]
	}
	isGreaterRight := dfs(node.Right, nums, res)
	if isGreaterRight {
		*nums = (*nums)[:len(*nums)-1]
	}
	return isGreater
}

func GoodNodes(root *TreeNode) int {
	maxNums := []int{math.MinInt}
	res := 0
	dfs(root, &maxNums, &res)
	return res
}
