package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}
	levelNodes := [][]*TreeNode{}
	res = append(res, []int{root.Val})
	levelNodes = append(levelNodes, []*TreeNode{root})
	i := 0
	for i < len(levelNodes) {
		nodes := levelNodes[i]
		newNodes := []*TreeNode{}
		newResItem := []int{}
		for _, node := range nodes {
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
				newResItem = append(newResItem, node.Left.Val)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
				newResItem = append(newResItem, node.Right.Val)
			}
		}
		if len(newNodes) > 0 {
			levelNodes = append(levelNodes, newNodes)
			res = append(res, newResItem)
		}
		i++
	}
	return res
}
