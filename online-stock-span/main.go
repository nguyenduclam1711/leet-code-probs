package onlinestockspan

type StockSpanner struct {
	stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	res := 1
	for len(this.stack) > 0 {
		lastItem := this.stack[len(this.stack)-1]
		if price < lastItem[1] {
			break
		}
		res += lastItem[0]
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, [2]int{res, price})
	return res
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
