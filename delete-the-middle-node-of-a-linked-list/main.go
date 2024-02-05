package deletethemiddlenodeofalinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteMiddle(head *ListNode) *ListNode {
	nodes := []*ListNode{}
	curr := head
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	if len(nodes) == 1 {
		return nil
	}

	middle := len(nodes) / 2
	if middle == len(nodes)-1 {
		nodes[middle-1].Next = nil
	} else {
		nodes[middle-1].Next = nodes[middle].Next
	}
	return nodes[0]
}
