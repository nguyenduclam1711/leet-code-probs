/**
 * @param {number[]} nums
 * @param {number} k
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var rotate = function (nums, k) {
  const newArr = Array(nums.length);
  for (let i = 0; i < nums.length; i++) {
    const newPos = (i + k) % nums.length;
    newArr[newPos] = nums[i];
  }
  for (let i = 0; i < newArr.length; i++) {
    nums[i] = newArr[i];
  }
};
