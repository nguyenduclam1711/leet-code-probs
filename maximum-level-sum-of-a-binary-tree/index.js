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
 * @return {number}
 */
var maxLevelSum = function (root) {
  let q = [[root]];
  let maxSum = -Infinity;
  let level = 0;
  let res = 0;

  while (q[level].length > 0) {
    q.push([]);
    let sum = 0;

    for (const n of q[level]) {
      sum += n.val;
      if (n.left) {
        q[level + 1].push(n.left);
      }
      if (n.right) {
        q[level + 1].push(n.right);
      }
    }
    level++;
    if (sum > maxSum) {
      maxSum = sum;
      res = level;
    }
  }
  return res;
};
