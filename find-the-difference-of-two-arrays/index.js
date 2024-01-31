/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[][]}
 */
var findDifference = function(nums1, nums2) {
  const map1 = nums1.reduce((acc, n) => {
    acc[n] = true;
    return acc;
  }, {});
  const map2 = nums2.reduce((acc, n) => {
    acc[n] = true;
    return acc;
  }, {});
  const res = [[], []];

  for (const n of Object.keys(map1)) {
    if (!map2[n]) {
      res[0].push(n);
    }
  }
  for (const n of Object.keys(map2)) {
    if (!map1[n]) {
      res[1].push(n);
    }
  }
  return res;
};
