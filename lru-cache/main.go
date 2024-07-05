package lrucache

type Node struct {
	Val  int
	Key  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	capacity     int
	currCapacity int
	headNode     *Node
	tailNode     *Node
	mapNodes     map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:     capacity,
		currCapacity: 0,
		mapNodes:     map[int]*Node{},
	}
}

func (this *LRUCache) SwapToHead(node *Node) {
	if this.headNode != node {
		if node.Prev != nil {
			node.Prev.Next = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}
		node.Next = this.headNode
		if this.tailNode == node {
			this.tailNode = node.Prev
		}
		node.Prev = nil
		this.headNode.Prev = node
		this.headNode = node
	}
}

func (this *LRUCache) AddToHead(key int, value int) *Node {
	newNode := &Node{
		Key: key,
		Val: value,
	}
	if this.headNode == nil {
		this.headNode = newNode
		this.tailNode = newNode
	} else {
		this.headNode.Prev = newNode
		newNode.Next = this.headNode
		this.headNode = newNode
	}
	this.mapNodes[key] = newNode
	return newNode
}

func (this *LRUCache) RemoveTail() {
	if this.tailNode == nil {
		return
	}
	key := this.tailNode.Key
	delete(this.mapNodes, key)
	if this.tailNode.Prev != nil {
		this.tailNode.Prev.Next = nil
	}
	if this.tailNode == this.headNode {
		this.headNode = nil
	}
	this.tailNode = this.tailNode.Prev
}

func (this *LRUCache) Get(key int) int {
	node, e := this.mapNodes[key]

	if !e {
		return -1
	}
	this.SwapToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, exist := this.mapNodes[key]

	if exist {
		node.Val = value
		this.SwapToHead(node)
	} else {
		if this.currCapacity < this.capacity {
			this.AddToHead(key, value)
			this.currCapacity++
		} else {
			this.RemoveTail()
			this.AddToHead(key, value)
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
