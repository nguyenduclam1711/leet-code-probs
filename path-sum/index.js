function recurse(node, targetSum, prevSum) {
  if (!node) {
    return false;
  }
  const currSum = node.val + prevSum;

  if (!node.left && !node.right && currSum === targetSum) {
    return true;
  }
  return (
    recurse(node.left, targetSum, currSum) ||
    recurse(node.right, targetSum, currSum)
  );
}

/**
 * @param {TreeNode} root
 * @param {number} targetSum
 * @return {boolean}
 */
var hasPathSum = function(root, targetSum) {
  return recurse(root, targetSum, 0);
};
