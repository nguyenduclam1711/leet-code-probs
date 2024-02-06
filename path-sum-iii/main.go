package pathsumiii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeVal(nums *[]int, val int) {
	if len(*nums) == 0 {
		return
	}
	*nums = (*nums)[:len(*nums)-1]
	for i := range *nums {
		(*nums)[i] -= val
	}
}

func dfs(node *TreeNode, targetSum int, pathNums *[]int, res *int) {
	if node == nil {
		return
	}
	if node.Val == targetSum {
		*res++
	}
	for i := range *pathNums {
		(*pathNums)[i] += node.Val
		if (*pathNums)[i] == targetSum {
			*res++
		}
	}
	(*pathNums) = append((*pathNums), node.Val)

	dfs(node.Left, targetSum, pathNums, res)
	dfs(node.Right, targetSum, pathNums, res)
	removeVal(pathNums, node.Val)
}

func PathSum(root *TreeNode, targetSum int) int {
	pathNums := []int{}
	res := 0
	dfs(root, targetSum, &pathNums, &res)
	return res
}
