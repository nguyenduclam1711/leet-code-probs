/**
 * @param {number} num
 * @param {number} k
 * @param {number} n
 * @param {number} total
 * @param {number[]} nums
 * @param {number[][]} res
 * @return {void}
 */
function dfs(num, k, n, total, nums, res) {
  if (total > n) {
    return;
  }
  if (nums.length === k) {
    if (total === n) {
      res.push([...nums]);
    }
    return;
  }
  for (let i = num + 1; i < 10; i++) {
    nums.push(i);
    dfs(i, k, n, total + i, nums, res);
    nums.pop();
  }
}

/**
 * @param {number} k
 * @param {number} n
 * @return {number[][]}
 */
var combinationSum3 = function(k, n) {
  const res = [];
  const nums = [];
  dfs(0, k, n, 0, nums, res);
  return res;
};
