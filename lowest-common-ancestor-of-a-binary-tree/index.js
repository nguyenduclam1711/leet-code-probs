function dfs(node, p, q, trackingPath, paths) {
  if (!node) {
    return;
  }
  trackingPath.push(node);
  if (node == p) {
    paths.path1 = [...trackingPath];
  }
  if (node == q) {
    paths.path2 = [...trackingPath];
  }
  if (paths.path1.length == 0 || paths.path2.length == 0) {
    dfs(node.left, targetNode, trackingPath, paths, pathName);
    dfs(node.right, targetNode, trackingPath, paths, pathName);
  }
  trackingPath.pop();
}

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
var lowestCommonAncestor = function(root, p, q) {
  const paths = {};
  dfs(root, p, q, [], paths);

  let p1 = paths.path1.length - 1;
  let p2 = paths.path2.length - 1;
  while (p1 >= 0 && p2 >= 0) {
    if (p1 > p2) {
      p1--;
    } else if (p2 > p1) {
      p2--;
    } else {
      if (paths.path1[p1] === paths.path2[p2]) {
        return paths.path1[p1];
      }
      p1--;
      p2--;
    }
  }
  return p;
};
