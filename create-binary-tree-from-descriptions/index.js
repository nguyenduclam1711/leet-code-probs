/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number[][]} descriptions
 * @return {TreeNode}
 */
var createBinaryTree = function(descriptions) {
  const mapNode = {};
  const mapChild = {};

  function createNode(val) {
    if (typeof mapNode[val] !== "undefined") {
      return mapNode[val];
    }
    mapNode[val] = new TreeNode(val);
    return mapNode[val];
  }

  for (let i = 0; i < descriptions.length; i++) {
    const desc = descriptions[i];
    const parentNode = createNode(desc[0]);
    const childNode = createNode(desc[1]);
    mapChild[desc[1]] = true;

    if (desc[2] === 1) {
      parentNode.left = childNode;
    } else {
      parentNode.right = childNode;
    }
  }
  let head = null;
  const allVals = Object.keys(mapNode);

  for (const val of allVals) {
    if (!mapChild[val]) {
      head = mapNode[val];
      break;
    }
  }
  return head;
};
