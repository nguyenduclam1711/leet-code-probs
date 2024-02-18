/**
 * @param {number[]} nums
 * @return {number}
 */
var findPeakElement = function(nums) {
  let l = 0;
  let r = nums.length;
  while (l < r) {
    const m = Math.floor((l + r) / 2);
    const adjL = m - 1;
    const adjR = m + 1;
    const isValidL = adjL < 0 || nums[adjL] < nums[m];
    const isValidR = adjR > nums.length - 1 || nums[adjR] < nums[m];
    if (isValidL && isValidR) {
      return m;
    } else if (!isValidL) {
      r = m;
    } else {
      l = m + 1;
    }
  }
  return -1;
};
