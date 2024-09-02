/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var sortList = function(head) {
  if (!head) {
    return head;
  }
  const arr = [];
  let curr = head;

  while (curr) {
    arr.push(curr);
    curr = curr.next;
  }
  arr.sort((node1, node2) => node1.val - node2.val);
  for (let i = 0; i < arr.length; i++) {
    const node = arr[i];
    if (i < arr.length - 1) {
      node.next = arr[i + 1];
    } else {
      node.next = null;
    }
  }
  return arr[0];
};
