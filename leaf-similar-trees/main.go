package leafsimilartrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, s *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*s = append(*s, node.Val)
		return
	}
	dfs(node.Left, s)
	dfs(node.Right, s)
}

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	s1 := []int{}
	s2 := []int{}
	dfs(root1, &s1)
	dfs(root2, &s2)
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
