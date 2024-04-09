var StockSpanner = function () {
  this.stack = [];
};

/**
 * @param {number} price
 * @return {number}
 */
StockSpanner.prototype.next = function (price) {
  let res = 1;

  while (this.stack.length > 0) {
    const lastItem = this.stack[this.stack.length - 1];
    if (price < lastItem[1]) {
      break;
    }
    res += lastItem[0];
    this.stack.pop();
  }
  this.stack.push([res, price]);
  return res;
};

/**
 * Your StockSpanner object will be instantiated and called as such:
 * var obj = new StockSpanner()
 * var param_1 = obj.next(price)
 */
