/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersection = function(nums1, nums2) {
  const result = []
  const uniqNums = {}

  for (const num of nums1) {
    if (!uniqNums[num]) {
      uniqNums[num] = true
    }
  }
  for (const num of nums2) {
    if (uniqNums[num]) {
      result.push(num)
      delete uniqNums[num]
    }
  }
  return result
};

console.log(intersection([1,2,2,1], [2,2]))
