package partitionlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	var leftHead, leftTail, rightHead, rightTail, curr *ListNode
	curr = head

	for curr != nil {
		nextCurr := curr.Next
		if curr.Val < x {
			if leftHead == nil {
				leftHead = curr
			}
			if leftTail != nil {
				leftTail.Next = curr
			}
			leftTail = curr
		} else {
			if rightHead == nil {
				rightHead = curr
			}
			if rightTail != nil {
				rightTail.Next = curr
			}
			rightTail = curr
		}
		curr = nextCurr
	}
	if leftTail != nil {
		leftTail.Next = rightHead

		if rightTail != nil {
			rightTail.Next = nil
		}
		return leftHead
	}
	if rightTail != nil {
		rightTail.Next = nil
	}
	return rightHead
}
