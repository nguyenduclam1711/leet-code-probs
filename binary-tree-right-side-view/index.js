/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var rightSideView = function(root) {
  const res = [];

  if (!root) {
    return res;
  }
  let q = [root];

  while (q.length > 0) {
    res.push(q[q.length - 1].val);
    const newQ = [];
    for (const n of q) {
      if (n.left) {
        newQ.push(n.left);
      }
      if (n.right) {
        newQ.push(n.right);
      }
    }
    q = newQ;
  }
  return res;
};
