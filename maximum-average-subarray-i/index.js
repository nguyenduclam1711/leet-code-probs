/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findMaxAverage = function(nums, k) {
  let i = 0;
  let j = 0;
  let currTotal = 0;
  let totalMax = -Infinity;

  while (j < nums.length) {
    let length = j - i + 1;
    currTotal += nums[j];

    if (length > k) {
      currTotal -= nums[i];
      i++;
      length--;
    }
    if (length === k && currTotal > totalMax) {
      totalMax = currTotal;
    }
    j++;
  }
  return totalMax / k;
};

console.log(findMaxAverage([1, 12, -5, -6, 50, 3], 4));
