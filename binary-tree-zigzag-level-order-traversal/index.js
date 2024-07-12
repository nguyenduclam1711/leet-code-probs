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
 * @return {number[][]}
 */
var zigzagLevelOrder = function (root) {
  if (!root) {
    return [];
  }
  const res = [];
  let isLeft = true;
  let q = [root];

  while (q.length > 0) {
    const newQ = [];
    res.push([]);

    for (let i = 0; i < q.length; i++) {
      const n = q[i];
      if (n.left) {
        newQ.push(n.left);
      }
      if (n.right) {
        newQ.push(n.right);
      }
      if (isLeft) {
        res[res.length - 1].push(n.val);
      } else {
        res[res.length - 1].unshift(n.val);
      }
    }
    isLeft = !isLeft;
    q = newQ;
  }
  return res;
};
