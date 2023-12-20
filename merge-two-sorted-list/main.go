package mergetwosortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1 := list1
	p2 := list2
	list := LinkedList{}

	for p1 != nil || p2 != nil {
		if p1 != nil && p2 != nil {
			if p1.Val <= p2.Val {
				list.insert(p1.Val)
				p1 = p1.Next
			} else {
				list.insert(p2.Val)
				p2 = p2.Next
			}
		} else if p1 == nil && p2 != nil {
			list.insert(p2.Val)
			p2 = p2.Next
		} else {
			list.insert(p1.Val)
			p1 = p1.Next
		}
	}
	return list.head
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func (l *LinkedList) insert(val int) {
	newNode := ListNode{
		Val: val,
	}
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
	} else {
		l.tail.Next = &newNode
		l.tail = &newNode
	}
}
