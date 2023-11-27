/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
  const hashMap = {};
  let i = 0;

  while (i < nums.length) {
    if (hashMap[target - nums[i]] !== undefined) {
      return [i, hashMap[target - nums[i]]];
    }
    hashMap[nums[i]] = i;
    i++;
  }
  return [];
};

console.log(`Input: nums = [2,7,11,15], target = 9`);
console.log(`Output ${twoSum([2, 7, 11, 15], 9)}`);
