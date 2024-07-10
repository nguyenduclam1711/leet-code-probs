/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
function recurse(node, obj) {
  if (!node) {
    return;
  }
  obj.count++;
  recurse(node.left, obj);
  recurse(node.right, obj);
}
/**
 * @param {TreeNode} root
 * @return {number}
 */
var countNodes = function(root) {
  const obj = {
    count: 0,
  };
  recurse(root, obj);
  return obj.count;
};
