/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function(nums, val) {
  let count = 0

  for (const num of nums) {
    if (num !== val) {
      nums[count] = num
      count++
    }
  }
  return count
};

console.log(removeElement([3,2,2,3], 3))
