package numberofrecentcalls

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	low := t - 3000
	for i, time := range this.queue {
		if time >= low {
			this.queue = this.queue[i:]
			break
		}
	}
	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
