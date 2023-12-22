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
 * @return {number[][]}
 */
var levelOrder = function (root) {
  const res = [];

  if (!root) {
    return res;
  }
  const levelNodes = [[root]];
  res.push([root.val]);
  let i = 0;

  while (i < levelNodes.length) {
    const nodes = levelNodes[i];
    const newNodes = [];
    const newResItem = [];

    for (const node of nodes) {
      if (node.left) {
        newNodes.push(node.left);
        newResItem.push(node.left.val);
      }
      if (node.right) {
        newNodes.push(node.right);
        newResItem.push(node.right.val);
      }
    }
    if (newNodes.length > 0) {
      levelNodes.push(newNodes);
      res.push(newResItem);
    }
    i++;
  }
  return res;
};
