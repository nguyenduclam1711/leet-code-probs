package reorderlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReorderList(head *ListNode) {
	s := []*ListNode{}
	curr := head

	for curr != nil {
		s = append(s, curr)
		curr = curr.Next
	}
	l, r := 0, len(s)-1
	m := (l + r) / 2
	for l <= r {
		s[l].Next = s[r]
		if l == m {
			s[r].Next = nil
		} else {
			s[r].Next = s[l+1]
		}
		l++
		r--
	}
}
