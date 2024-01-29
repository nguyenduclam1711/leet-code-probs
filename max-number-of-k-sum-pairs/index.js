/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxOperations = function(nums, k) {
  const sorted = nums.sort((a, b) => a - b);
  let l = 0;
  let r = sorted.length - 1;
  let res = 0;

  while (l < r) {
    const dif = k - sorted[l];

    if (dif === sorted[r]) {
      res++;
      l++;
      r--;
    } else if (dif > sorted[r]) {
      l++;
    } else {
      r--;
    }
  }
  return res;
};
