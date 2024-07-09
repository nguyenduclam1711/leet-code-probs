package flattenbinarytreetolinkedlist

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var head, currLinkedListNode *TreeNode
	var recurseFunc func(currNode *TreeNode)

	recurseFunc = func(currNode *TreeNode) {
		if currNode == nil {
			return
		}
		newNode := &TreeNode{
			Val: currNode.Val,
		}
		if currLinkedListNode == nil {
			head = newNode
		} else {
			currLinkedListNode.Right = newNode
		}
		currLinkedListNode = newNode
		recurseFunc(currNode.Left)
		recurseFunc(currNode.Right)
	}
	recurseFunc(root)
	root.Right = head.Right
	root.Left = nil
}
