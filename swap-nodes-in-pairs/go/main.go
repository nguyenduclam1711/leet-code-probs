package swappairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var prev *ListNode = nil
	currNode := head
	newHead := head

	for currNode != nil && currNode.Next != nil {
		nextNode := currNode.Next
		currNode.Next = nextNode.Next
		nextNode.Next = currNode
		if prev != nil {
			prev.Next = nextNode
		}
		currNode = currNode.Next
		prev = nextNode.Next
		if newHead == head {
			newHead = nextNode
		}
	}
	return newHead
}
