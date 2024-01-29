/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxOperations = function(nums, k) {
  // using 2 pointers
  const sorted = nums.sort((a, b) => a - b);
  let l = 0;
  let r = sorted.length - 1;
  let res = 0;

  while (l < r) {
    const dif = k - sorted[l];

    if (dif === sorted[r]) {
      res++;
      l++;
      r--;
    } else if (dif > sorted[r]) {
      l++;
    } else {
      r--;
    }
  }
  return res;
};

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxOperations2 = function(nums, k) {
  // using hashmap
  const hashmap = {};
  let res = 0;

  for (const n of nums) {
    if (hashmap[n] > 0) {
      hashmap[n]--;
      res++;
    } else {
      const dif = k - n;
      if (dif in hashmap) {
        hashmap[dif]++;
      } else {
        hashmap[dif] = 1;
      }
    }
  }
  return res;
};
