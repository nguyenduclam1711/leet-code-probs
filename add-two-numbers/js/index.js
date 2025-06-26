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
var addTwoNumbers = function (l1, l2) {
  let p1 = l1;
  let p2 = l2;
  let head = null;
  let l = head;
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

    if (!head) {
      head = new ListNode(newVal);
      l = head;
    } else {
      l.next = new ListNode(newVal);
      l = l.next;
    }
  }
  if (needToAdded === 1) {
    l.next = new ListNode(1);
  }
  return head;
};
