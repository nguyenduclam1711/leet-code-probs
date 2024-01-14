package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	s := map[int]*Node{}
	queue := []*Node{node}
	visited := map[int]bool{}

	for len(queue) > 0 {
		curr := queue[0]
		if !visited[curr.Val] {
			visited[curr.Val] = true
			if s[curr.Val] == nil {
				s[curr.Val] = &Node{
					Val: curr.Val,
				}
			}
			for _, n := range curr.Neighbors {
				if s[n.Val] == nil {
					s[n.Val] = &Node{
						Val: n.Val,
					}
				}
				s[curr.Val].Neighbors = append(s[curr.Val].Neighbors, s[n.Val])
				queue = append(queue, n)
			}
		}
		queue = queue[1:]
	}
	return s[1]
}
