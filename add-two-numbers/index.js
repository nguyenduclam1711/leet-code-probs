class ListNode {
  constructor(val, next) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
  p1 = l1;
  p2 = l2;
  l = new LinkedList();
  let needToAdded = 0;

  while (p1 || p2) {
    let val1;
    let val2;
    if (p1) {
      val1 = p1.val;
      p1 = p1.next;
    } else {
      val1 = 0;
    }
    if (p2) {
      val2 = p2.val;
      p2 = p2.next;
    } else {
      val2 = 0;
    }
    let newVal = val1 + val2 + needToAdded;

    if (newVal < 10) {
      needToAdded = 0;
    } else {
      newVal = newVal % 10;
      needToAdded = 1;
    }
    l.insert(newVal);
  }
  if (needToAdded === 1) {
    l.insert(1);
  }
  return l.head;
};

class LinkedList {
  constructor() {
    this.head = this.tail = undefined;
  }

  insert(val) {
    const newNode = new ListNode(val);
    if (!this.head) {
      this.head = this.tail = newNode;
      return;
    }
    this.tail.next = newNode;
    this.tail = newNode;
  }
}
