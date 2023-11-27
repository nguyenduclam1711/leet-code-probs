/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function (nums) {
  const appearTimes = nums.length / 2;
  const map = {};

  for (const num of nums) {
    if (num in map) {
      map[num]++;
    } else {
      map[num] = 1;
    }
    if (map[num] > appearTimes) {
      return num;
    }
  }
  return -1;
};

console.log(majorityElement([2, 2, 1, 1, 1, 2, 2]));
