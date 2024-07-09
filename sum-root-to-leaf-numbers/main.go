package sumroottoleafnumbers

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recurse(currNode *TreeNode, prevStr string, res *int) {
	if currNode == nil {
		return
	}
	currStr := fmt.Sprint(prevStr, currNode.Val)
	if currNode.Left == nil && currNode.Right == nil {
		currNum, _ := strconv.Atoi(currStr)

		*res += currNum
		return
	}
	recurse(currNode.Left, currStr, res)
	recurse(currNode.Right, currStr, res)
}

func sumNumbers(root *TreeNode) int {
	res := 0
	recurse(root, "", &res)
	return res
}
