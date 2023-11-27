/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function (nums) {
  const storeObj = {};
  let i = 0;

  while (i < nums.length) {
    const num = nums[i];
    if (storeObj[num]) {
      return true;
    }
    storeObj[num] = true;
    i++;
  }
  return false;
};

console.log("Input: nums = [1,2,3,1]");
console.log("Output:", containsDuplicate([1, 2, 3, 1]));
