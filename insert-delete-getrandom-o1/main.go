package insertdeletegetrandomo1

import "math/rand"

type RandomizedSet struct {
	data         []int
	mapDataToPos map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		data:         []int{},
		mapDataToPos: map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, e := this.mapDataToPos[val]; e {
		return false
	}
	this.data = append(this.data, val)
	this.mapDataToPos[val] = len(this.data) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	curPos, e := this.mapDataToPos[val]
	if !e {
		return false
	}
	lastPos := len(this.data) - 1
	lastVal := this.data[len(this.data)-1]
	this.data[curPos], this.data[lastPos] = this.data[lastPos], this.data[curPos]
	this.data = this.data[:len(this.data)-1]
	this.mapDataToPos[lastVal] = curPos
	delete(this.mapDataToPos, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	randPos := rand.Intn(len(this.data))
	return this.data[randPos]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
