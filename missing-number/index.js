/**
 * @param {number[]} nums
 * @return {number}
 */
var missingNumber = function(nums) {
  const totalInNumsShouldBe = (nums.length * (nums.length + 1)) / 2;
  let totalInNums = 0;

  for (const num of nums) {
    totalInNums += num;
  }
  return totalInNumsShouldBe - totalInNums;
};

console.log(missingNumber([0, 1, 3]));
