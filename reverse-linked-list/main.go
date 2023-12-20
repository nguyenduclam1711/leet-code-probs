package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	var newHead *ListNode

	for curr != nil {
		newHead = &ListNode{
			Val:  curr.Val,
			Next: newHead,
		}
		curr = curr.Next
	}
	return newHead
}
