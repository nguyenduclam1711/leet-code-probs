/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersect = function (nums1, nums2) {
  const result = [];
  const uniqNums = {};

  for (const num of nums1) {
    if (uniqNums[num]) {
      uniqNums[num]++;
    } else {
      uniqNums[num] = 1;
    }
  }
  for (const num of nums2) {
    if (uniqNums[num] > 0) {
      result.push(num);
      uniqNums[num]--;
    }
  }
  return result;
};

console.log(intersect([1, 2, 2, 1, 3, 3, 4, 5, 1], [2, 2, 1, 1, 1, 3, 4, 5]));
