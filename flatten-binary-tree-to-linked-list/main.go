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
	head := &TreeNode{
		Val: root.Val,
	}
	currLinkedListNode := head

	var recurseFunc func(currNode *TreeNode)

	recurseFunc = func(currNode *TreeNode) {
		if currNode == nil {
			return
		}
		newNode := &TreeNode{
			Val: currNode.Val,
		}
		currLinkedListNode.Right = newNode
		currLinkedListNode = newNode
		recurseFunc(currNode.Left)
		recurseFunc(currNode.Right)
	}

	recurseFunc(root.Left)
	recurseFunc(root.Right)
	root.Right = head.Right
	root.Left = nil
}
