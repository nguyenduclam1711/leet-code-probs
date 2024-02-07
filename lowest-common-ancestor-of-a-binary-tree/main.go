package lowestcommonancestorofabinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node, p, q *TreeNode, trackingPath, path1, path2 *[]*TreeNode) {
	if node == nil {
		return
	}
	*trackingPath = append(*trackingPath, node)
	if node == p {
		*path1 = append([]*TreeNode{}, *trackingPath...)
	}
	if node == q {
		*path2 = append([]*TreeNode{}, *trackingPath...)
	}
	if len(*path1) == 0 || len(*path2) == 0 {
		dfs(node.Left, p, q, trackingPath, path1, path2)
		dfs(node.Right, p, q, trackingPath, path1, path2)
	}
	*trackingPath = (*trackingPath)[:len(*trackingPath)-1]
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var path1, path2, trackingPath []*TreeNode
	dfs(root, p, q, &trackingPath, &path1, &path2)

	p1, p2 := len(path1)-1, len(path2)-1
	for p1 >= 0 && p2 >= 0 {
		if p1 > p2 {
			p1--
		} else if p2 > p1 {
			p2--
		} else {
			if path1[p1] == path2[p2] {
				return path1[p1]
			}
			p1--
			p2--
		}
	}
	return p
}
