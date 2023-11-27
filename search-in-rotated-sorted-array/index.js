/**
 * @param {number[]} nums
 * @return {number}
 */
var findMinIndex = function (nums) {
  let l = 0;
  let r = nums.length - 1;

  while (l < r) {
    const mdl = parseInt((l + r) / 2);

    if (nums[mdl] > nums[r]) {
      l = mdl + 1;
    } else {
      r = mdl;
    }
  }
  return l;
};

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function (nums, target) {
  let minIndex = findMinIndex(nums);
  let l = 0;
  let r = nums.length - 1;

  if (target === nums[minIndex]) {
    return minIndex;
  }
  if (target > nums[minIndex] && target <= nums[r]) {
    l = minIndex;
  } else {
    r = minIndex - 1;
  }
  while (l < r) {
    const mdl = parseInt((l + r) / 2);
    if (target > nums[mdl]) {
      l = mdl + 1;
    } else {
      r = mdl;
    }
  }
  if (nums[l] === target) {
    return l;
  }
  return -1;
};

console.log(search([4, 5, 6, 7, 0, 1, 2], 0));
