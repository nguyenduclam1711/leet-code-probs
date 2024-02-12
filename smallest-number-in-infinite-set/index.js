class SmallestInfiniteSet {
  constructor() {
    this.data = [];
    this.set = Array(1001).fill(false);

    for (let i = 0; i < 1000; i++) {
      this.addBack(i + 1);
    }
  }

  /**
   * @return {number}
   */
  popSmallest() {
    const pop = this.data[0];
    this.data[0] = this.data[this.data.length - 1];
    this.data[this.data.length - 1] = pop;
    this.data.pop();
    this.heapifyDown(0);
    this.set[pop] = false;
    return pop;
  }

  /**
   * @param {number} num
   * @return {void}
   */
  addBack(num) {
    if (this.set[num]) {
      return;
    }
    this.data.push(num);
    this.heapifyUp(this.data.length - 1);
    this.set[num] = true;
  }

  heapifyDown(i) {
    const l = lChild(i);
    if (l >= this.data.length) {
      return;
    }
    const r = rChild(i);
    if (r >= this.data.length) {
      if (this.data[i] > this.data[l]) {
        const tmp = this.data[i];
        this.data[i] = this.data[l];
        this.data[l] = tmp;
        this.heapifyDown(l);
      }
      return;
    }
    if (this.data[l] < this.data[r]) {
      if (this.data[i] > this.data[l]) {
        const tmp = this.data[i];
        this.data[i] = this.data[l];
        this.data[l] = tmp;
        this.heapifyDown(l);
      }
    } else {
      if (this.data[i] > this.data[r]) {
        const tmp = this.data[i];
        this.data[i] = this.data[r];
        this.data[r] = tmp;
        this.heapifyDown(r);
      }
    }
  }

  heapifyUp(i) {
    const p = parent(i);
    if (p < 0) {
      return;
    }
    if (this.data[i] < this.data[p]) {
      const tmp = this.data[i];
      this.data[i] = this.data[p];
      this.data[p] = tmp;
      this.heapifyUp(p);
    }
  }
}

function parent(i) {
  return Math.floor((i - 1) / 2);
}

function lChild(i) {
  return i * 2 + 1;
}

function rChild(i) {
  return i * 2 + 2;
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * var obj = new SmallestInfiniteSet()
 * var param_1 = obj.popSmallest()
 * obj.addBack(num)
 */
