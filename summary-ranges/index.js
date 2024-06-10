/**
 * @param {number[]} nums
 * @return {string[]}
 */
var summaryRanges = function(nums) {
  if (nums.length === 0) {
    return [];
  }
  let low = nums[0];
  let hi = nums[0];
  const res = [];
  for (let i = 1; i < nums.length; i++) {
    const n = nums[i];
    if (hi + 1 === n) {
      hi = n;
      continue;
    }
    if (low === hi) {
      res.push(low.toString());
    } else {
      res.push(`${low}->${hi}`);
    }
    low = n;
    hi = n;
  }
  if (low === hi) {
    res.push(low.toString());
  } else {
    res.push(`${low}->${hi}`);
  }
  return res;
};
