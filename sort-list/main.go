package sortlist

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	s := []*ListNode{}
	curr := head

	for curr != nil {
		s = append(s, curr)
		curr = curr.Next
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].Val < s[j].Val
	})

	for i, node := range s {
		if i < len(s)-1 {
			node.Next = s[i+1]
		} else {
			node.Next = nil
		}
	}
	return s[0]
}
