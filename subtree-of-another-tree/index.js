function checkSame(root, subRoot) {
  if (!root && !subRoot) {
    return true;
  }
  if (!root || !subRoot) {
    return false;
  }
  if (root.val !== subRoot.val) {
    return false;
  }
  return (
    checkSame(root.left, subRoot.left) && checkSame(root.right, subRoot.right)
  );
}

function walk(root, subRoot) {
  if (!root) {
    return false;
  }
  if (root.val === subRoot.val) {
    const isSame = checkSame(root, subRoot);
    if (isSame) {
      return isSame;
    }
  }
  return walk(root.right, subRoot) || walk(root.left, subRoot);
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
 * @param {TreeNode} subRoot
 * @return {boolean}
 */
var isSubtree = function (root, subRoot) {
  return walk(root, subRoot);
};
