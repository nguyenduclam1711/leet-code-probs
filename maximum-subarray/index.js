/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function (nums) {
  const numsLength = nums.length;
  const prefix = Array(numsLength);
  let max = 0;

  for (let i = 0; i < numsLength; i++) {
    const num = nums[i];
    if (i === 0) {
      prefix[i] = num;
      max = prefix[i];
      continue;
    }
    const newNum = num + prefix[i - 1];
    if (newNum > num) {
      prefix[i] = newNum;
    } else {
      prefix[i] = num;
    }
    if (prefix[i] > max) {
      max = prefix[i];
    }
  }
  return max;
};

console.log(maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4]));
