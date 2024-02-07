package binarytreerightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RightSideView(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return res
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		currLen := len(q)
		if q[currLen-1] != nil {
			res = append(res, q[currLen-1].Val)
		}
		newQ := []*TreeNode{}
		for _, n := range q {
			if n.Left != nil {
				newQ = append(newQ, n.Left)
			}
			if n.Right != nil {
				newQ = append(newQ, n.Right)
			}
		}
		q = newQ
	}
	return res
}
