package minstack

import "math"

type Node struct {
	next *Node
	prev *Node
	val  int
}

type MinStack struct {
	head    *Node
	tail    *Node
	mins    []int
	currMin int
}

func Constructor() MinStack {
	return MinStack{
		currMin: math.MaxInt,
	}
}

func (this *MinStack) Push(val int) {
	if val < this.currMin {
		this.currMin = val
	}
	this.mins = append(this.mins, this.currMin)

	newNode := &Node{
		val: val,
	}
	if this.head == nil {
		this.head = newNode
		this.tail = newNode
		return
	}
	this.tail.next = newNode
	newNode.prev = this.tail
	this.tail = newNode
}

func (this *MinStack) Pop() {
	if this.tail == nil {
		return
	}
	this.mins = this.mins[:len(this.mins)-1]
	if len(this.mins) > 0 {
		this.currMin = this.mins[len(this.mins)-1]
	} else {
		this.currMin = math.MaxInt
	}
	if this.tail.prev != nil {
		this.tail.prev.next = nil
		this.tail = this.tail.prev
	}
}

func (this *MinStack) Top() int {
	if this.tail == nil {
		return 0
	}
	return this.tail.val
}

func (this *MinStack) GetMin() int {
	return this.currMin
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
