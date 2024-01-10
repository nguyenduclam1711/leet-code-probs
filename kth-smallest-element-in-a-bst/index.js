function dfs(node, arr, k) {
  if (!node) {
    return;
  }
  dfs(node.left, arr, k);
  arr.push(node.val);
  if (arr.length === k) {
    return;
  }
  dfs(node.right, arr, k);
}

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
 * @param {number} k
 * @return {number}
 */
var kthSmallest = function (root, k) {
  const arr = [];
  dfs(root, arr, k);
  return arr[k - 1];
};
