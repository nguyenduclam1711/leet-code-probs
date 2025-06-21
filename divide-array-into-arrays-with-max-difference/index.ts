function divideArray(nums: number[], k: number): number[][] {
  const res: number[][] = [];
  nums.sort((a, b) => a - b);
  let p1 = 0;

  while (p1 < nums.length) {
    if (nums[p1 + 2] - nums[p1] <= k) {
      res.push(nums.slice(p1, p1 + 3));
      p1 += 3;
    } else {
      return [];
    }
  }
  return res;
};
