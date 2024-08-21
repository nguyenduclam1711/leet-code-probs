package binarysearchtreeiterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	currentIndex int
	inorderArr   []int
}

func Constructor(root *TreeNode) BSTIterator {
	inorderArr := []int{}
	var addToInorderArr func(curr *TreeNode)
	addToInorderArr = func(curr *TreeNode) {
		if curr == nil {
			return
		}
		addToInorderArr(curr.Left)
		inorderArr = append(inorderArr, curr.Val)
		addToInorderArr(curr.Right)
	}
	addToInorderArr(root)
	return BSTIterator{
		currentIndex: 0,
		inorderArr:   inorderArr,
	}
}

func (this *BSTIterator) Next() int {
	res := this.inorderArr[this.currentIndex]
	this.currentIndex++
	return res
}

func (this *BSTIterator) HasNext() bool {
	return this.currentIndex < len(this.inorderArr)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
