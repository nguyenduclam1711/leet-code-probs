var RandomizedSet = function() {
  this.data = [];
  this.mapDataToPos = {};
};

/**
 * @param {number} val
 * @return {boolean}
 */
RandomizedSet.prototype.insert = function(val) {
  if (this.mapDataToPos[val] !== undefined) {
    return false;
  }
  this.data.push(val);
  this.mapDataToPos[val] = this.data.length - 1;
  return true;
};

/**
 * @param {number} val
 * @return {boolean}
 */
RandomizedSet.prototype.remove = function(val) {
  const curPos = this.mapDataToPos[val];
  if (curPos === undefined) {
    return false;
  }
  const lastPos = this.data.length - 1;
  const lastVal = this.data[lastPos];
  this.data[lastPos] = val;
  this.data[curPos] = lastVal;
  this.data.pop();
  this.mapDataToPos[lastVal] = curPos;
  delete this.mapDataToPos[val];
  return true;
};

/**
 * @return {number}
 */
RandomizedSet.prototype.getRandom = function() {
  const randPos = Math.floor(Math.random() * this.data.length);
  return this.data[randPos];
};

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * var obj = new RandomizedSet()
 * var param_1 = obj.insert(val)
 * var param_2 = obj.remove(val)
 * var param_3 = obj.getRandom()
 */
