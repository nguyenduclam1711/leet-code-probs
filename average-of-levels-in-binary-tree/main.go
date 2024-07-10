package averageoflevelsinbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		newQ := []*TreeNode{}
		total := 0

		for _, node := range q {
			if node.Left != nil {
				newQ = append(newQ, node.Left)
			}
			if node.Right != nil {
				newQ = append(newQ, node.Right)
			}
			total += node.Val
		}
		res = append(res, float64(total)/float64(len(q)))
		q = newQ
	}
	return res
}
