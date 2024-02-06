package longestzigzagpathinabinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, prevDirection byte, prevLength int, res *int) {
	if node == nil {
		return
	}
	if prevLength > *res {
		*res = prevLength
	}
	if prevDirection == 'l' {
		dfs(node.Left, 'l', 1, res)
		dfs(node.Right, 'r', prevLength+1, res)
	} else {
		dfs(node.Left, 'l', prevLength+1, res)
		dfs(node.Right, 'r', 1, res)
	}
}

func LongestZigZag(root *TreeNode) int {
	res := 0
	dfs(root, 'l', 0, &res)
	return res
}
