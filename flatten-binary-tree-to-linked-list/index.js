function recurse(currNode, obj) {
  if (!currNode) {
    return;
  }
  const newNode = new TreeNode(currNode.val);
  obj.curr.right = newNode;
  obj.curr = newNode;

  recurse(currNode.left, obj);
  recurse(currNode.right, obj);
}
/**
 * @param {TreeNode} root
 * @return {void} Do not return anything, modify root in-place instead.
 */
var flatten = function(root) {
  if (!root) {
    return;
  }
  const head = new TreeNode(root);
  const obj = {
    curr: head,
  };
  recurse(root.left, obj);
  recurse(root.right, obj);
  root.left = null;
  root.right = head.right;
};
