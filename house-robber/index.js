function findMax(nums) {
  return Math.max(...nums);
}

/**
 * @param {number[]} nums
 * @return {number}
 */
var rob = function (nums) {
  max = Array(nums.length);

  for (let i = nums.length - 1; i >= 0; i--) {
    if (i === nums.length - 1 || i === nums.length - 2) {
      max[i] = nums[i];
    } else {
      max[i] = nums[i] + findMax(max.slice(i + 2));
    }
  }
  if (max.length === 1) {
    return max[0];
  }
  if (max[0] > max[1]) {
    return max[0];
  }
  return max[1];
};
