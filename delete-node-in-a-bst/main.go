package deletenodeinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteLargestNode(node *TreeNode, prevNode *TreeNode) *TreeNode {
	if node.Right == nil {
		if prevNode.Right == node {
			prevNode.Right = node.Left
		} else {
			prevNode.Left = node.Left
		}
		return node
	}
	return deleteLargestNode(node.Right, node)
}

func searchAndDelete(node *TreeNode, prevNode *TreeNode, key int) {
	if node == nil {
		return
	}
	if node.Val < key {
		searchAndDelete(node.Right, node, key)
	} else if node.Val > key {
		searchAndDelete(node.Left, node, key)
	} else {
		if node.Left == nil && node.Right == nil {
			// node doesnt have childs
			if prevNode.Right == node {
				prevNode.Right = nil
			} else {
				prevNode.Left = nil
			}
		} else if node.Left != nil && node.Right != nil {
			// node have 2 childs
			// get the largest node on the left of current node and replace the current node with the largest node on the left
			replaceNode := deleteLargestNode(node.Left, node)
			replaceNode.Right = node.Right
			if node.Left != nil {
				replaceNode.Left = node.Left
			}
			if prevNode.Left == node {
				prevNode.Left = replaceNode
			} else {
				prevNode.Right = replaceNode
			}
		} else {
			// node have 1 child
			if node.Left != nil {
				if prevNode.Right == node {
					prevNode.Right = node.Left
				} else {
					prevNode.Left = node.Left
				}
			} else {
				if prevNode.Right == node {
					prevNode.Right = node.Right
				} else {
					prevNode.Left = node.Right
				}
			}
		}
	}
}

func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	head := &TreeNode{Left: root}
	searchAndDelete(root, head, key)
	return head.Left
}
