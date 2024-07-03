package removeduplicatesfromsortedlistii

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var newHead, prevNewHead, curr, currNumNode *ListNode
	count := 0
	currNumNode = head
	curr = head

	for curr != nil {
		if curr.Val == currNumNode.Val {
			count++
		} else {
			if count == 1 {
				if newHead == nil {
					newHead = currNumNode
				}
				if prevNewHead != nil {
					prevNewHead.Next = currNumNode
				}
				prevNewHead = currNumNode
			}
			count = 1
			currNumNode = curr
		}
		curr = curr.Next
	}
	if count == 1 {
		if newHead == nil {
			newHead = currNumNode
		}
		if prevNewHead != nil {
			prevNewHead.Next = currNumNode
		}
		prevNewHead = currNumNode
	}
	if prevNewHead != nil {
		prevNewHead.Next = nil
	}
	return newHead
}
