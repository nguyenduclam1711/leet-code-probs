/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */

function recurse(currNode, preStr, obj) {
  if (!currNode) {
    return;
  }
  const currStr = `${preStr}${currNode.val}`;
  if (!currNode.left && !currNode.right) {
    const currNum = Number(currStr);
    obj.res += currNum;
    return;
  }
  recurse(currNode.left, currStr, obj);
  recurse(currNode.right, currStr, obj);
}

/**
 * @param {TreeNode} root
 * @return {number}
 */
var sumNumbers = function(root) {
  const obj = {
    res: 0,
  };
  recurse(root, "", obj);
  return obj.res;
};
