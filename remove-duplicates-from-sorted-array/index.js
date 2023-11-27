/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
  let count = 1;
  let prevNum = nums[0];

  for (let i = 1; i < nums.length; i++) {
    if (nums[i] !== prevNum) {
      nums[count] = nums[i];
      count++;
      prevNum = nums[i];
    }
  }
  return count;
};

console.log(removeDuplicates([1, 1, 2]));
