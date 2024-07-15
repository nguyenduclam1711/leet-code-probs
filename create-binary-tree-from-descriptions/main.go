package createbinarytreefromdescriptions

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	mapNode := map[int]*TreeNode{}
	mapChild := map[int]bool{}

	createNode := func(val int) *TreeNode {
		if node, e := mapNode[val]; e {
			return node
		}
		mapNode[val] = &TreeNode{
			Val: val,
		}
		return mapNode[val]
	}

	for _, desc := range descriptions {
		parentVal, childVal, isLeft := desc[0], desc[1], desc[2]
		parentNode := createNode(parentVal)
		childNode := createNode(childVal)
		mapChild[childVal] = true

		if isLeft == 1 {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}
	}
	var head *TreeNode
	for val, node := range mapNode {
		if !mapChild[val] {
			head = node
			break
		}
	}
	return head
}
