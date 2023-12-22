function recurse(root) {
  if (!root) {
    return;
  }
  const tmp = root.left;
  root.left = root.right;
  root.right = tmp;

  recurse(root.left);
  recurse(root.right);
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
 * @return {TreeNode}
 */
var invertTree = function (root) {
  recurse(root);
  return root;
};
