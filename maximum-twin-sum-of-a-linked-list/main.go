package maximumtwinsumofalinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

func PairSum(head *ListNode) int {
	length := 0
	curr := head
	for curr != nil {
		curr = curr.Next
		length++
	}

	middle := length / 2
	res := 0
	p1 := head
	p2 := head
	for i := 0; i < middle; i++ {
		p2 = p2.Next
	}
	p2 = reverseLinkedList(p2)
	for p2 != nil {
		sum := p1.Val + p2.Val
		if sum > res {
			res = sum
		}
		p2 = p2.Next
		p1 = p1.Next
	}
	return res
}
