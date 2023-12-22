function walk(node, min, max) {
  if (node.val > max) {
    return walk(node.left, min, max);
  } else if (node.val < min) {
    return walk(node.right, min, max);
  } else {
    return node;
  }
}

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
var lowestCommonAncestor = function (root, p, q) {
  let max = p.val;
  let min = q.val;
  if (q.val > p.val) {
    max = q.val;
    min = p.val;
  }
  return walk(root, min, max);
};
