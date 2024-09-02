/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permute = function (nums) {
  const visitedPos = {};
  const arr = [];
  const res = [];

  function recurse(pos) {
    if (visitedPos[pos]) {
      return;
    }
    visitedPos[pos] = true;
    arr.push(nums[pos]);

    if (arr.length === nums.length) {
      res.push([...arr]);
    } else {
      for (let i = 0; i < nums.length; i++) {
        recurse(i);
      }
    }
    visitedPos[pos] = false;
    arr.pop();
  }

  for (let i = 0; i < nums.length; i++) {
    recurse(i);
  }
  return res;
};
