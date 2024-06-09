/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var containsNearbyDuplicate = function (nums, k) {
  const m = new Map();
  for (let i = 0; i < nums.length; i++) {
    const n = nums[i];

    if (!m.has(n)) {
      m.set(n, i);
    } else {
      const pos = m.get(n);
      const diffAbs = Math.abs(i - pos);
      if (diffAbs <= k) {
        return true;
      }
      m.set(n, i);
    }
  }
  return false;
};
