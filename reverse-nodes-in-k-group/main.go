package reversenodesinkgroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(node *ListNode, prevGroupNode *ListNode, left int, right int) (*ListNode, *ListNode) {
	i := left
	var groupHead, groupTail, curr *ListNode
	curr = node

	for i >= left && i <= right {
		nextCurr := curr.Next
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
		if prevGroupNode != nil {
			prevGroupNode.Next = groupHead
		}
		curr = nextCurr
		i++
	}
	return groupHead, groupTail
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	curr := head
	listLength := 0
	for curr != nil {
		listLength++
		curr = curr.Next
	}

	curr = head
	var res, prev *ListNode
	for i := 1; i <= listLength && curr != nil; i += k {
		leftLength := listLength - i + 1
		if k <= leftLength {
			groupHead, groupTail := reverseBetween(curr, prev, i, i+k-1)
			curr = groupTail.Next
			prev = groupTail

			if i == 1 {
				res = groupHead
			}
		}
	}
	return res
}
