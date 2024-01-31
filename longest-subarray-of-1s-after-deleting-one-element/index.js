/**
 * @param {number[]} nums
 * @return {number}
 */
var longestSubarray = function(nums) {
  let l = 0;
  let r = 0;
  let numberOfZeroes = 0;
  let res = 0;

  while (r < nums.length) {
    if (nums[r] === 0) {
      numberOfZeroes++;
    }
    if (numberOfZeroes <= 1) {
      const windowLen = r - l;
      if (windowLen > res) {
        res = windowLen;
      }
    } else {
      if (nums[l] === 0) {
        numberOfZeroes--;
      }
      l++;
    }
    r++;
  }
  return res;
};
