package reverselinkedlistii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	i := 1
	var groupHead, groupTail, prevGroupHead, curr *ListNode
	curr = head

	for curr != nil {
		nextCurr := curr.Next
		if i == left-1 {
			prevGroupHead = curr
		} else if i >= left && i <= right {
			if i == left {
				groupHead = curr
			} else if i == left+1 {
				groupTail = curr
				groupHead.Next = groupTail.Next
				groupTail.Next = groupHead
				groupHead, groupTail = groupTail, groupHead
			} else {
				groupTail.Next = curr.Next
				curr.Next = groupHead
				groupHead = curr
			}
			if prevGroupHead != nil {
				prevGroupHead.Next = groupHead
			}
		}
		curr = nextCurr
		i++
	}
	if left == 1 {
		return groupHead
	}
	return head
}
