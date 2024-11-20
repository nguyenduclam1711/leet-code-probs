/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maximumSubarraySum = function(nums, k) {
  let l = 0;
  let r = k - 1;
  let res = 0;
  const currNums = new Map();
  let currSum = 0;

  for (let i = l; i <= r; i++) {
    currNums.set(nums[i], (currNums.get(nums[i]) ?? 0) + 1);
    currSum += nums[i];
  }
  if (currNums.size === k && currSum > res) {
    res = currSum;
  }

  while (r < nums.length) {
    currSum -= nums[l];
    currNums.set(nums[l], currNums.get(nums[l]) - 1);
    if (!currNums.get(nums[l])) {
      currNums.delete(nums[l]);
    }
    l++;
    r++;
    if (r === nums.length) {
      break;
    }
    currNums.set(nums[r], (currNums.get(nums[r]) ?? 0) + 1);
    currSum += nums[r];
    if (currNums.size === k && currSum > res) {
      res = currSum;
    }
  }
  return res;
};
