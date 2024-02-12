package smallestnumberininfiniteset

type SmallestInfiniteSet struct {
	data []int
	set  []bool
}

func Constructor() SmallestInfiniteSet {
	res := SmallestInfiniteSet{
		data: []int{},
		set:  make([]bool, 1001),
	}
	for i := 0; i < 1000; i++ {
		res.AddBack(i + 1)
	}
	return res
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	pop := this.data[0]
	this.data[0], this.data[len(this.data)-1] = this.data[len(this.data)-1], this.data[0]
	this.data = this.data[:len(this.data)-1]
	this.heapifyDown(0)
	this.set[pop] = false
	return pop
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if this.set[num] {
		return
	}
	this.data = append(this.data, num)
	this.heapifyUp(len(this.data) - 1)
	this.set[num] = true
}

func (this *SmallestInfiniteSet) parentIdx(i int) int {
	return (i - 1) / 2
}

func (this *SmallestInfiniteSet) lChildIdx(i int) int {
	return i*2 + 1
}

func (this *SmallestInfiniteSet) rChildIdx(i int) int {
	return i*2 + 2
}

func (this *SmallestInfiniteSet) heapifyUp(i int) {
	p := this.parentIdx(i)

	if p < 0 {
		return
	}
	if this.data[i] < this.data[p] {
		this.data[p], this.data[i] = this.data[i], this.data[p]
		this.heapifyUp(p)
	}
}

func (this *SmallestInfiniteSet) heapifyDown(i int) {
	l := this.lChildIdx(i)
	if l >= len(this.data) {
		return
	}
	lC := this.data[l]
	r := this.rChildIdx(i)
	if r >= len(this.data) {
		if this.data[i] > this.data[l] {
			this.data[i], this.data[l] = this.data[l], this.data[i]
			this.heapifyDown(l)
		}
		return
	}
	rC := this.data[r]

	if lC < rC {
		if this.data[i] > this.data[l] {
			this.data[i], this.data[l] = this.data[l], this.data[i]
			this.heapifyDown(l)
		}
	} else {
		if this.data[i] > this.data[r] {
			this.data[i], this.data[r] = this.data[r], this.data[i]
			this.heapifyDown(r)
		}
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
