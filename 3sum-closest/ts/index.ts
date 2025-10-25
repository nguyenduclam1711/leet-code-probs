function threeSumClosest(nums: number[], target: number): number {
  nums.sort((a, b) => a - b);
  let minDiff = Infinity;
  let result = 0;

  for (let i = 0; i < nums.length; i++) {
    let l = i + 1;
    let r = nums.length - 1;

    while (l < r) {
      const currSum = nums[i] + nums[l] + nums[r];
      if (currSum === target) {
        return currSum;
      }
      const diff = Math.abs(target - currSum);
      if (diff < minDiff) {
        minDiff = diff;
        result = currSum;
      }
      if (currSum < target) {
        l++;
      } else {
        r--;
      }
    }
  }
  return result;
};
