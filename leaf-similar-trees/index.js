function dfs(node, arr) {
  if (!node) {
    return;
  }
  if (!node.left && !node.right) {
    arr.push(node.val);
    return;
  }
  dfs(node.left, arr);
  dfs(node.right, arr);
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
 * @param {TreeNode} root1
 * @param {TreeNode} root2
 * @return {boolean}
 */
var leafSimilar = function(root1, root2) {
  const s1 = [];
  const s2 = [];
  dfs(root1, s1);
  dfs(root2, s2);
  if (s1.length != s2.length) {
    return false;
  }
  for (let i = 0; i < s1.length; i++) {
    if (s1[i] != s2[i]) {
      return false;
    }
  }
  return true;
};
