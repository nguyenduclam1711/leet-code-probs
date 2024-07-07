package constructbinarytreefrominorderandpostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	res := &TreeNode{
		Val: rootVal,
	}

	var rootPosInInorder int
	for i := range inorder {
		if inorder[i] == rootVal {
			rootPosInInorder = i
			break
		}
	}

	if rootPosInInorder == 0 && len(inorder) > 1 {
		// only have right node
		res.Right = buildTree(inorder[rootPosInInorder+1:], postorder[:len(postorder)-1])
	} else if rootPosInInorder == len(inorder)-1 && len(inorder) > 1 {
		// only have left node
		res.Left = buildTree(inorder[:rootPosInInorder], postorder[:len(postorder)-1])
	} else if len(inorder) > 2 {
		// have both right and left node
		rightLen := len(inorder) - rootPosInInorder
		var rightPost int
		for i := range postorder {
			if len(postorder)-i == rightLen {
				rightPost = i
				break
			}
		}
		res.Left = buildTree(inorder[:rootPosInInorder], postorder[:rightPost])
		res.Right = buildTree(inorder[rootPosInInorder+1:], postorder[rightPost:len(postorder)-1])
	}
	return res
}
