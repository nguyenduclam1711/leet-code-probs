/**
 * @param {number[]} nums
 * @return {number}
 */
var pivotIndex = function(nums) {
  const leftSum = Array(nums.length);
  const rightSum = Array(nums.length);

  for (let i = 0; i < nums.length; i++) {
    if (i === 0) {
      leftSum[i] = 0;
    } else {
      leftSum[i] = leftSum[i - 1] + nums[i - 1];
    }
  }
  for (let i = nums.length - 1; i >= 0; i--) {
    if (i === nums.length - 1) {
      rightSum[i] = 0;
    } else {
      rightSum[i] = rightSum[i + 1] + nums[i + 1];
    }
  }
  for (let i = 0; i < nums.length; i++) {
    if (leftSum[i] === rightSum[i]) {
      return i;
    }
  }
  return -1;
};
