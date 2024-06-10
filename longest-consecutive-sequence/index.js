/**
 * @param {number} curr
 * @param {{ [num: number]: boolean }} appears
 * @param {{ [num: number]: number }} cached
 * @param {number} currLen
 * @return {number}
 */
function recurse(curr, appears, cached, currLen) {
  if (typeof cached[curr] !== "undefined") {
    return cached[curr];
  }
  const nextN = curr + 1;
  let res = currLen;

  if (appears[nextN]) {
    res += recurse(nextN, appears, cached, 1);
  }
  cached[curr] = res;
  return res;
}

/**
 * @param {number[]} nums
 * @return {number}
 */
var longestConsecutive = function(nums) {
  const appears = {};
  for (let i = 0; i < nums.length; i++) {
    appears[nums[i]] = true;
  }

  // store the length of consecutive by each num
  const cached = {};
  let res = 0;

  for (let i = 0; i < nums.length; i++) {
    const currRes = recurse(nums[i], appears, cached, 1);
    if (currRes > res) {
      res = currRes;
    }
  }
  return res;
};
