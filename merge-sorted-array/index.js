/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function(nums1, m, nums2, n) {
  const sortedArr = [...nums1.slice(0, m), ...nums2.slice(0, n)].sort((a, b) => a - b)
  for (let i = 0; i < sortedArr.length; i++) {
    nums1[i] = sortedArr[i]
  }
};

const nums1 = [1,2,3,0,0,0]

console.log(merge(nums1, 3, [2,5,6], 3))
console.log(nums1)
