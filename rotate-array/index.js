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

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var rotate2 = function (nums, k) {
  const reverse = (start, end) => {
    while (start < end) {
      const tmp = nums[end];
      nums[end] = nums[start];
      nums[start] = tmp;
      start++;
      end--;
    }
  };
  const j = k % nums.length;
  reverse(0, nums.length - 1);
  reverse(0, j - 1);
  reverse(j, nums.length - 1);
};
