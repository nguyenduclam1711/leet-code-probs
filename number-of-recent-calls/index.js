var RecentCounter = function () {
  this.queue = [];
};

/**
 * @param {number} t
 * @return {number}
 */
RecentCounter.prototype.ping = function (t) {
  this.queue.push(t);
  const low = t - 3000;
  for (let i = 0; i < this.queue.length; i++) {
    const time = this.queue[i];
    if (time >= low) {
      this.queue = this.queue.slice(i);
      break;
    }
  }
  return this.queue.length;
};

/**
 * Your RecentCounter object will be instantiated and called as such:
 * var obj = new RecentCounter()
 * var param_1 = obj.ping(t)
 */
