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
var maxPathSum = function (root) {
  let res = -Infinity;

  const recurse = function (node) {
    if (!node) {
      return 0;
    }
    if (!node.left && !node.right) {
      if (node.val > res) {
        res = node.val;
      }
      return node.val;
    }
    const maxLeft = recurse(node.left);
    const maxRight = recurse(node.right);
    const totalPathRight = node.val + maxRight;
    const totalPathLeft = node.val + maxLeft;
    const currNodeMax = Math.max(node.val, totalPathLeft, totalPathRight);

    if (currNodeMax === node.val) {
      if (currNodeMax > res) {
        res = currNodeMax;
      }
    } else if (currNodeMax === totalPathLeft) {
      const newRes = Math.max(currNodeMax, currNodeMax + maxRight);
      if (newRes > res) {
        res = newRes;
      }
    } else if (currNodeMax === totalPathRight) {
      const newRes = Math.max(currNodeMax, currNodeMax + maxLeft);
      if (newRes > res) {
        res = newRes;
      }
    }
    return currNodeMax;
  };

  recurse(root);
  return res;
};
