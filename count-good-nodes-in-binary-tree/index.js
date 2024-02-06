/**
 * @param {TreeNode} node
 * @param {Array<number>} maxNums
 * @param {Array<number>} result
 * @return {void}
 */
function dfs(node, maxNums, result) {
  if (!node) {
    return result;
  }
  let isAppendToNums = false;
  if (node.val >= maxNums[maxNums.length - 1]) {
    isAppendToNums = true;
    result[0]++;
    maxNums.push(node.val);
  }
  dfs(node.left, maxNums, result);
  dfs(node.right, maxNums, result);
  if (isAppendToNums) {
    maxNums.pop();
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
var goodNodes = function(root) {
  const maxNums = [-Infinity];
  const result = [0];
  dfs(root, maxNums, result);
  return result[0];
};
