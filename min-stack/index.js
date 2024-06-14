var MinStack = function() {
  this.currMin = Infinity;
  this.mins = [];
  this.head = null;
  this.tail = null;
};

/**
 * @param {number} val
 * @return {void}
 */
MinStack.prototype.push = function(val) {
  if (val < this.currMin) {
    this.currMin = val;
  }
  this.mins.push(this.currMin);
  const newNode = {
    val,
  };

  if (!this.head) {
    this.head = this.tail = newNode;
    return;
  }
  this.tail.next = newNode;
  newNode.prev = this.tail;
  this.tail = newNode;
};

/**
 * @return {void}
 */
MinStack.prototype.pop = function() {
  if (!this.tail) {
    return;
  }
  this.mins.pop();
  if (this.mins.length > 0) {
    this.currMin = this.mins[this.mins.length - 1];
  } else {
    this.currMin = Infinity;
  }
  if (this.tail.prev) {
    this.tail.prev.next = null;
    this.tail = this.tail.prev;
  }
};

/**
 * @return {number}
 */
MinStack.prototype.top = function() {
  if (!this.tail) {
    return 0;
  }
  return this.tail.val;
};

/**
 * @return {number}
 */
MinStack.prototype.getMin = function() {
  return this.currMin;
};

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(val)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */
