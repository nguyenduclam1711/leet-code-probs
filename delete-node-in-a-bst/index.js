function deleteLargestNode(node, prevNode) {
  if (!node.right) {
    if (prevNode.right == node) {
      prevNode.right = node.left;
    } else {
      prevNode.left = node.left;
    }
    return node;
  }
  return deleteLargestNode(node.right, node);
}

function searchAndDelete(node, prevNode, key) {
  if (!node) {
    return;
  }
  if (node.val < key) {
    searchAndDelete(node.right, node, key);
  } else if (node.val > key) {
    searchAndDelete(node.left, node, key);
  } else {
    if (!node.left && !node.right) {
      // node doesn't have childs
      if (prevNode.right == node) {
        prevNode.right = null;
      } else {
        prevNode.left = null;
      }
    } else if (node.left && node.right) {
      // node have 2 childs
      const replaceNode = deleteLargestNode(node.left, node);
      replaceNode.right = node.right;
      if (node.left) {
        replaceNode.left = node.left;
      }
      if (prevNode.left == node) {
        prevNode.left = replaceNode;
      } else {
        prevNode.right = replaceNode;
      }
    } else {
      // node have 1 child
      if (node.left) {
        if (prevNode.right == node) {
          prevNode.right = node.left;
        } else {
          prevNode.left = node.left;
        }
      } else {
        if (prevNode.right == node) {
          prevNode.right = node.right;
        } else {
          prevNode.left = node.right;
        }
      }
    }
  }
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
 * @param {number} key
 * @return {TreeNode}
 */
var deleteNode = function(root, key) {
  if (!root) {
    return null;
  }
  const head = new TreeNode(0, root);
  searchAndDelete(root, head, key);
  return head.left;
};
