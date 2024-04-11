/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
  let apps = 1;
  let curr = nums[0];
  let pos = 1;

  for (let i = 1; i < nums.length; i++) {
    if (curr === nums[i]) {
      if (apps < 2) {
        apps++;
        if (pos < i) {
          nums[pos] = nums[i];
        }
        pos++;
      }
    } else {
      curr = nums[i];
      apps = 1;
      if (pos < i) {
        nums[pos] = nums[i];
      }
      pos++;
    }
  }
  return pos;
};
