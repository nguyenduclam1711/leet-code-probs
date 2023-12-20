package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	m := map[int]*ListNode{}
	length := 0
	curr := head

	for curr != nil {
		m[length] = curr
		length++
		curr = curr.Next
	}
	removeIdx := length - n
	if removeIdx == 0 {
		m[0] = m[1]
	} else {
		m[removeIdx-1].Next = m[removeIdx+1]
		m[removeIdx] = nil
	}
	return m[0]
}
