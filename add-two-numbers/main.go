package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	l := LinkedList{}
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
		l.insert(newVal)
	}
	if needToAdded == 1 {
		l.insert(1)
	}
	return l.head
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func (l *LinkedList) insert(value int) {
	newNode := &ListNode{
		Val: value,
	}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	l.tail.Next = newNode
	l.tail = newNode
}
