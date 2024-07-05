package populatingnextrightpointersineachnodeii

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	q := []*Node{root}

	for len(q) > 0 {
		nextQ := []*Node{}
		for i := range q {
			if i < len(q)-1 {
				q[i].Next = q[i+1]
			}
			if q[i].Left != nil {
				nextQ = append(nextQ, q[i].Left)
			}
			if q[i].Right != nil {
				nextQ = append(nextQ, q[i].Right)
			}
		}
		q = nextQ
	}
	return root
}
