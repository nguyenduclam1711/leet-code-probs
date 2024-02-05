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
var deleteMiddle = function (head) {
  const nodes = [];
  let curr = head;
  while (curr) {
    nodes.push(curr);
    curr = curr.next;
  }

  if (nodes.length === 1) {
    return null;
  }

  const middle = Math.floor(nodes.length / 2);
  if (middle === nodes.length - 1) {
    nodes[middle - 1].next = null;
  } else {
    nodes[middle - 1].next = nodes[middle].next;
  }
  return nodes[0];
};
