/**
 * @param {number} target
 * @param {number[]} nums
 * @return {number}
 */
var minSubArrayLen = function(target, nums) {
  let windowLength = nums.length + 1;
  let p1 = 0;
  let p2 = 0;
  let currSum = 0 - nums[p1];

  while (p1 <= p2 && p1 < nums.length && p2 < nums.length) {
    currSum += nums[p1] + nums[p2];
    if (currSum < target) {
      currSum -= nums[p1];
      p2++;
    } else {
      currLength = p2 - p1 + 1;
      if (currLength < windowLength) {
        windowLength = currLength;
      }
      if (currSum == target) {
        currSum -= nums[p1];
        p1++;
        if (p1 > p2 || p1 >= nums.length) {
          break;
        }
        currSum -= nums[p1];
        p2++;
      } else {
        currSum -= nums[p1];
        p1++;
        if (p1 > p2 || p1 >= nums.length) {
          break;
        }
        currSum -= nums[p1];
        currSum -= nums[p2];
      }
    }
  }
  if (windowLength === nums.length + 1) {
    return 0;
  }
  return windowLength;
};
