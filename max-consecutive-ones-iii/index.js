/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var longestOnes = function(nums, k) {
  let l = 0;
  let r = 0;
  let res = 0;
  let numberOfZeroes = 0;

  while (r < nums.length) {
    if (nums[r] === 0) {
      numberOfZeroes++;
    }
    if (numberOfZeroes > k) {
      if (nums[l] === 0) {
        numberOfZeroes--;
      }
      l++;
    }
    const windowLen = r - l + 1;
    if (windowLen > res) {
      res = windowLen;
    }
    r++;
  }
  return res;
};
