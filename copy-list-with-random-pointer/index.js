/**
 * // Definition for a _Node.
 * function _Node(val, next, random) {
 *    this.val = val;
 *    this.next = next;
 *    this.random = random;
 * };
 */

/**
 * @param {_Node} head
 * @return {_Node}
 */
var copyRandomList = function(head) {
  let currHead = null;
  let currNode = null;
  const mapping = new Map();

  currNode = head;
  while (currNode) {
    const newNode = new _Node(currNode.val, null, null);
    if (!currHead) {
      currHead = newNode;
    }
    mapping.set(currNode, newNode);
    currNode = currNode.next;
  }

  currNode = head;
  while (currNode) {
    if (mapping.has(currNode)) {
      if (mapping.has(currNode.next)) {
        mapping.get(currNode).next = mapping.get(currNode.next);
      }
      if (mapping.has(currNode.random)) {
        mapping.get(currNode).random = mapping.get(currNode.random);
      }
    }
    currNode = currNode.next;
  }
  return currHead;
};
