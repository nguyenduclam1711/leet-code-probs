/**
 * @param {TreeNode} node
 * @param {string} prevDirection
 * @param {number} prevLen
 * @param {Array<number>} result
 * @return {void}
 */
function dfs(node, prevDirection, prevLen, result) {
  if (!node) {
    return;
  }
  if (prevLen > result[0]) {
    result[0] = prevLen;
  }
  if (prevDirection == "l") {
    dfs(node.left, "l", 1, result);
    dfs(node.right, "r", prevLen + 1, result);
  } else {
    dfs(node.left, "l", prevLen + 1, result);
    dfs(node.right, "r", 1, result);
  }
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
 * @return {number}
 */
var longestZigZag = function(root) {
  const result = [0];
  dfs(root, "l", 0, result);
  return result[0];
};
