/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function (nums) {
  const s = new Set();
  for (const n of nums) {
    if (s.has(n)) {
      return true;
    }
    s.add(n);
  }
  return false;
};

console.log("Input: nums = [1,2,3,1]");
console.log("Output:", containsDuplicate([1, 2, 3, 1]));
