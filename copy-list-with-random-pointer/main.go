package copylistwithrandompointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	var currHead, currNode *Node

	currNode = head
	for currNode != nil {
		newNode := Node{
			Val: currNode.Val,
		}
		if currHead == nil {
			currHead = &newNode
		}
		m[currNode] = &newNode
		currNode = currNode.Next
	}

	currNode = head
	for currNode != nil {
		node, exist := m[currNode]
		if exist {
			node.Next = m[currNode.Next]
			node.Random = m[currNode.Random]
		}
		currNode = currNode.Next
	}
	return currHead
}
