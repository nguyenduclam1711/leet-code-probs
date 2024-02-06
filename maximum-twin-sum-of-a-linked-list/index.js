function reverse(head) {
  let prev = null;
  let curr = head;
  while (curr) {
    let tmp = curr.next;
    curr.next = prev;
    prev = curr;
    curr = tmp;
  }
  return prev;
}

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {number}
 */
var pairSum = function(head) {
  let curr = head;
  let length = 0;
  while (curr) {
    curr = curr.next;
    length++;
  }

  let res = 0;
  let p1 = head;
  let p2 = head;
  const middle = Math.floor(length / 2);
  for (let i = 0; i < middle; i++) {
    p2 = p2.next;
  }
  p2 = reverse(p2);
  while (p2) {
    const sum = p1.val + p2.val;
    if (sum > res) {
      res = sum;
    }
    p1 = p1.next;
    p2 = p2.next;
  }
  return res;
};
