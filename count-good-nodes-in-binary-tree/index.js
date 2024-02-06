/**
 * @param {TreeNode} node
 * @param {Array<number>} maxNums
 * @param {number} prevResult
 * @return {number}
 */
function dfs(node, maxNums, prevResult) {
  if (!node) {
    return prevResult;
  }
  let result = prevResult;
  if (node.val >= maxNums[maxNums.length - 1]) {
    result++;
    maxNums.push(node.val);
  }
  const resultLeft = dfs(node.left, maxNums, result);
  if (resultLeft > result) {
    maxNums.pop();
  }
  const resultRight = dfs(node.right, maxNums, resultLeft);
  if (resultRight > resultLeft) {
    maxNums.pop();
  }
  return resultRight;
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
  return dfs(root, maxNums, 0);
};
