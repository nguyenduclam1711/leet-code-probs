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
var averageOfLevels = function(root) {
  let q = [root];
  let res = [];

  while (q.length > 0) {
    const newQ = [];
    let total = 0;

    for (let i = 0; i < q.length; i++) {
      const node = q[i];
      if (node.left) {
        newQ.push(node.left);
      }
      if (node.right) {
        newQ.push(node.right);
      }
      total += node.val;
    }
    res.push(total / q.length);
    q = newQ;
  }
  return res;
};
