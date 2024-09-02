function getMostLeftNode(node) {
  if (!node) {
    return node;
  }
  if (node.left) {
    return getMostLeftNode(node.left);
  }
  return node;
}

function getMostRightNode(node) {
  if (!node) {
    return node;
  }
  if (node.right) {
    return getMostRightNode(node.right);
  }
  return node;
}

function dfs(node, arr) {
  if (!node) {
    return;
  }
  const closestLeft = getMostRightNode(node.left);
  const closestRight = getMostLeftNode(node.right);

  if (closestLeft) {
    const newRes = node.val - closestLeft.val;
    if (newRes < arr[0]) {
      arr[0] = newRes;
    }
  }
  if (closestRight) {
    const newRes = closestRight.val - node.val;
    if (newRes < arr[0]) {
      arr[0] = newRes;
    }
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
 * @param {TreeNode} root
 * @return {number}
 */
var getMinimumDifference = function (root) {
  const arr = [Infinity];
  dfs(root, arr);
  return arr[0];
};
