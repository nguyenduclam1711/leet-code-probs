/**
 * @param {number[]} nums
 */
var threeSum = function (nums) {
  const sortedNums = nums.sort((a, b) => a - b);
  const result = [];
  let prevA = 0;

  for (let i = 0; i < sortedNums.length; i++) {
    const num = sortedNums[i];

    if (num === prevA && i !== 0) {
      continue;
    }
    let l = i + 1;
    let r = sortedNums.length - 1;
    // init prevB like this because this initial will never includes in nums (nums is sorted so I can do this)
    let prevB = sortedNums[0] - 1;

    while (l < r) {
      if (sortedNums[l] === prevB) {
        l++;
        continue;
      }
      let tmp = sortedNums[l] + sortedNums[r];

      if (tmp === -num) {
        result.push([num, sortedNums[l], sortedNums[r]]);
        prevB = sortedNums[l];
        l++;
        r--;
      } else if (tmp < -num) {
        l++;
      } else {
        r--;
      }
    }
    prevA = num;
  }
  return result;
};

console.log(threeSum([3, 0, -2, -1, 1, 2]));
