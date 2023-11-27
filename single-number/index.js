/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function (nums) {
  const map = {};

  for (const num of nums) {
    if (map[num]) {
      delete map[num];
    } else {
      map[num] = true;
    }
  }
  return Object.keys(map)[0];
};

console.log(singleNumber([4, 1, 2, 1, 2]));
