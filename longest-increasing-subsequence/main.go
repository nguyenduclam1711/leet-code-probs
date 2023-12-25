package longestincreasingsubsequence

type BTreeNode struct {
	MaxDist int
	Val     int
	Left    *BTreeNode
	Right   *BTreeNode
}

type BTree struct {
	Head    *BTreeNode
	MaxDist int
}

func insertRecurse(node *BTreeNode, val int, maxDist int) {
	if node == nil {
		return
	}
	if val > node.Val && node.Right == nil {
		node.Right = &BTreeNode{
			Val:     val,
			MaxDist: maxDist,
		}
		return
	}
	if val <= node.Val && node.Left == nil {
		node.Left = &BTreeNode{
			Val:     val,
			MaxDist: maxDist,
		}
		return
	}
	if val > node.Val {
		insertRecurse(node.Right, val, maxDist)
	} else {
		insertRecurse(node.Left, val, maxDist)
	}
}

func (b *BTree) insert(val int, maxDist int) {
	if maxDist > b.MaxDist {
		b.MaxDist = maxDist
	}
	if b.Head == nil {
		b.Head = &BTreeNode{
			Val:     val,
			MaxDist: maxDist,
		}
		return
	}
	insertRecurse(b.Head, val, maxDist)
}

func calcLongestSubsAtPos(node *BTreeNode, num int, max int) int {
	if node == nil {
		return max
	}
	if node.Val <= num {
		return calcLongestSubsAtPos(node.Right, num, max)
	}
	newDist := 1 + node.MaxDist
	if newDist > max {
		max = newDist
	}
	lMax := calcLongestSubsAtPos(node.Left, num, max)
	rMax := calcLongestSubsAtPos(node.Right, num, max)
	if lMax > rMax {
		return lMax
	}
	return rMax
}

func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tree := BTree{}
	tree.insert(nums[len(nums)-1], 1)
	for i := len(nums) - 2; i >= 0; i-- {
		maxDist := calcLongestSubsAtPos(tree.Head, nums[i], 1)
		tree.insert(nums[i], maxDist)
	}
	return tree.MaxDist
}
