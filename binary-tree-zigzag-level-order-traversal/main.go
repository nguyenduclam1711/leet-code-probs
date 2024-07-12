package binarytreezigzaglevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	isLeft := true
	q := []*TreeNode{root}

	for len(q) > 0 {
		newQ := []*TreeNode{}
		res = append(res, []int{})

		for _, n := range q {
			if n.Left != nil {
				newQ = append(newQ, n.Left)
			}
			if n.Right != nil {
				newQ = append(newQ, n.Right)
			}
			if isLeft {
				res[len(res)-1] = append(res[len(res)-1], n.Val)
			} else {
				res[len(res)-1] = append([]int{n.Val}, res[len(res)-1]...)
			}
		}
		if isLeft {
			isLeft = false
		} else {
			isLeft = true
		}
		q = newQ
	}
	return res
}
