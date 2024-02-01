package convertsortedarraytobinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := len(nums) / 2
	root := &TreeNode{
		Val:   nums[m],
		Left:  SortedArrayToBST(nums[:m]),
		Right: SortedArrayToBST(nums[m+1:]),
	}
	return root
}
