package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	var head *ListNode = nil
	l := head
	needToAdded := 0

	for p1 != nil || p2 != nil {
		var val1 int
		var val2 int
		if p1 != nil {
			val1 = p1.Val
			p1 = p1.Next
		} else {
			val1 = 0
		}
		if p2 != nil {
			val2 = p2.Val
			p2 = p2.Next
		} else {
			val2 = 0
		}
		newVal := val1 + val2 + needToAdded

		if newVal < 10 {
			needToAdded = 0
		} else {
			newVal = newVal % 10
			needToAdded = 1
		}
		if head == nil {
			head = &ListNode{
				Val: newVal,
			}
			l = head
		} else {
			l.Next = &ListNode{
				Val:  newVal,
				Next: nil,
			}
			l = l.Next
		}
	}
	if needToAdded == 1 {
		next := &ListNode{
			Val:  needToAdded,
			Next: nil,
		}
		l.Next = next
	}
	return head
}
