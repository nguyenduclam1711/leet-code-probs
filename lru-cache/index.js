class Node {
  key = 0;
  val = 0;
  next = null;
  prev = null;

  constructor(key, value) {
    this.key = key;
    this.val = value;
  }
}

class LRUCache {
  capacity = 0;
  headNode = null;
  tailNode = null;
  mapNodes = {};
  currCapacity = 0;

  constructor(capacity) {
    this.capacity = capacity;
  }

  swapToHead(node) {
    if (this.headNode !== node) {
      if (node.prev) {
        node.prev.next = node.next;
      }
      if (node.next) {
        node.next.prev = node.prev;
      }
      node.next = this.headNode;
      if (this.tailNode === node) {
        this.tailNode = node.prev;
      }
      node.prev = null;
      this.headNode.prev = node;
      this.headNode = node;
    }
  }

  addToHead(key, value) {
    const newNode = new Node(key, value);
    if (!this.headNode) {
      this.headNode = newNode;
      this.tailNode = newNode;
    } else {
      this.headNode.prev = newNode;
      newNode.next = this.headNode;
      this.headNode = newNode;
    }
    this.mapNodes[key] = newNode;
    return newNode;
  }

  removeTail() {
    if (!this.tailNode) {
      return;
    }
    const key = this.tailNode.key;
    delete this.mapNodes[key];
    if (this.tailNode.prev) {
      this.tailNode.prev.next = null;
    }
    if (this.tailNode === this.headNode) {
      this.headNode = null;
    }
    this.tailNode = this.tailNode.prev;
  }

  get(key) {
    const node = this.mapNodes[key];

    if (typeof node === "undefined") {
      return -1;
    }
    this.swapToHead(node);
    return node.val;
  }

  put(key, value) {
    const node = this.mapNodes[key];

    if (node) {
      node.val = value;
      this.swapToHead(node);
    } else {
      if (this.currCapacity < this.capacity) {
        this.addToHead(key, value);
        this.currCapacity++;
      } else {
        this.removeTail();
        this.addToHead(key, value);
      }
    }
  }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * var obj = new LRUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */
