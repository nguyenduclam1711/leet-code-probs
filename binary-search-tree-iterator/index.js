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
 */
var BSTIterator = function(root) {
  const inorderArr = [];
  const recurse = (curr) => {
    if (!curr) {
      return;
    }
    recurse(curr.left);
    inorderArr.push(curr.val);
    recurse(curr.right);
  };
  recurse(root);
  this.inorderArr = inorderArr;
  this.currentIndex = 0;
};

/**
 * @return {number}
 */
BSTIterator.prototype.next = function() {
  const res = this.inorderArr[this.currentIndex];
  this.currentIndex++;
  return res;
};

/**
 * @return {boolean}
 */
BSTIterator.prototype.hasNext = function() {
  return this.currentIndex < this.inorderArr.length;
};

/**
 * Your BSTIterator object will be instantiated and called as such:
 * var obj = new BSTIterator(root)
 * var param_1 = obj.next()
 * var param_2 = obj.hasNext()
 */
