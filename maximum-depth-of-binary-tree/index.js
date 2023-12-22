function dfs(node, depth) {
  if (!node) {
    return depth;
  }
  const depthL = dfs(node.left, depth + 1);
  const depthR = dfs(node.right, depth + 1);

  if (depthL > depthR) {
    return depthL;
  }
  return depthR;
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
var maxDepth = function (root) {
  return dfs(root, 0);
};
