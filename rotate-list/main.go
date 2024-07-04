package rotatelist

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	arr := []*ListNode{}
	curr := head

	for curr != nil {
		arr = append(arr, curr)
		curr = curr.Next
	}
	numberOfRotate := k % len(arr)

	if numberOfRotate == 0 {
		return head
	}
	lastPosInArr := len(arr) - numberOfRotate - 1
	arr[len(arr)-1].Next = head
	arr[lastPosInArr].Next = nil
	head = arr[lastPosInArr+1]
	return head
}
