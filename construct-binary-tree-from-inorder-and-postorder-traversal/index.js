/**
 * @param {number[]} inorder
 * @param {number[]} postorder
 * @return {TreeNode}
 */
var buildTree = function(inorder, postorder) {
  if (!postorder.length || !inorder.length) {
    return null;
  }
  const rootVal = postorder[postorder.length - 1];
  const res = new TreeNode(rootVal);

  let rootPosInInorder = -1;
  for (let i = 0; i < inorder.length; i++) {
    if (inorder[i] === rootVal) {
      rootPosInInorder = i;
      break;
    }
  }

  if (rootPosInInorder === 0 && inorder.length > 1) {
    // only have right node
    res.right = buildTree(
      inorder.slice(rootPosInInorder + 1),
      postorder.slice(0, postorder.length - 1),
    );
  } else if (rootPosInInorder === inorder.length - 1 && inorder.length > 1) {
    // only have left node
    res.left = buildTree(
      inorder.slice(0, rootPosInInorder),
      postorder.slice(0, postorder.length - 1),
    );
  } else if (inorder.length > 2) {
    // have both right and left node
    const rightLen = inorder.length - rootPosInInorder;
    let rightPost = -1;
    for (let i = 0; i < postorder.length; i++) {
      if (postorder.length - i === rightLen) {
        rightPost = i;
        break;
      }
    }
    res.left = buildTree(
      inorder.slice(0, rootPosInInorder),
      postorder.slice(0, rightPost),
    );
    res.Right = buildTree(
      inorder.slice(rootPosInInorder + 1),
      postorder.slice(rightPost, postorder.length - 1),
    );
  }
  return res;
};
