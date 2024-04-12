/**
 * @param {number[]} nums
 * @return {number}
 */
var jump = function(nums) {
  let i = 0;
  let res = 0;

  while (i < nums.length - 1) {
    res++;
    const n = nums[i];
    if (i + n >= nums.length - 1) {
      return res;
    }
    let maxJumpPos = -1;
    let nextPos = i;
    for (let j = i + 1; j <= i + n; j++) {
      const curJumpPos = nums[j] + j;
      if (curJumpPos > maxJumpPos) {
        maxJumpPos = curJumpPos;
        nextPos = j;
      }
    }
    i = nextPos;
  }
  return res;
};
