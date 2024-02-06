/**
 * @param {TreeNode} node
 * @param {number} targetSum
 * @param {Array<number>} pathNums
 * @param {Array<number>} res
 * @return {void}
 */
function dfs(node, targetSum, pathNums, res) {
  if (!node) {
    return;
  }
  if (node.val == targetSum) {
    res[0]++;
  }
  for (let i = 0; i < pathNums.length; i++) {
    pathNums[i] += node.val;
    if (pathNums[i] == targetSum) {
      res[0]++;
    }
  }
  pathNums.push(node.val);

  dfs(node.left, targetSum, pathNums, res);
  dfs(node.right, targetSum, pathNums, res);

  // remove val for backtracking
  pathNums.pop();
  for (let i = 0; i < pathNums.length; i++) {
    pathNums[i] -= node.val;
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
 * @param {number} targetSum
 * @return {number}
 */
var pathSum = function (root, targetSum) {
  const pathNums = [];
  const result = [0];
  dfs(root, targetSum, pathNums, result);
  return result[0];
};
