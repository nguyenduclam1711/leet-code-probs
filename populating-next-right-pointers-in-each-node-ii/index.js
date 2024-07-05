/**
 * // Definition for a _Node.
 * function _Node(val, left, right, next) {
 *    this.val = val === undefined ? null : val;
 *    this.left = left === undefined ? null : left;
 *    this.right = right === undefined ? null : right;
 *    this.next = next === undefined ? null : next;
 * };
 */

/**
 * @param {_Node} root
 * @return {_Node}
 */
var connect = function(root) {
  if (!root) {
    return root;
  }
  let q = [root];

  while (q.length > 0) {
    const nextQ = [];
    for (let i = 0; i < q.length; i++) {
      if (i < q.length - 1) {
        q[i].next = q[i + 1];
      }
      if (q[i].left) {
        nextQ.push(q[i].left);
      }
      if (q[i].right) {
        nextQ.push(q[i].right);
      }
    }
    q = nextQ;
  }
  return root;
};
