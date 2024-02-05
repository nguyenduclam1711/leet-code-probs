package oddevenlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	oddPointer := head
	evenPointer := head.Next
	firstEven := evenPointer

	for oddPointer != nil && evenPointer != nil {
		if evenPointer.Next == nil {
			break
		}
		oddPointer.Next = evenPointer.Next
		oddPointer = oddPointer.Next
		evenPointer.Next = oddPointer.Next
		evenPointer = evenPointer.Next
	}
	oddPointer.Next = firstEven
	return head
}
