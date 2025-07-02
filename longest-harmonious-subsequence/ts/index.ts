// counting
function findLHS1(nums: number[]): number {
  const map = new Map<string, number>()
  const uniqueNums = new Set<number>();
  for (const n of nums) {
    uniqueNums.add(n);
    const combination = `${n},${n + 1}`;
    map.set(combination, (map.get(combination) ?? 0) + 1);
    const combination2 = `${n - 1},${n}`;
    map.set(combination2, (map.get(combination2) ?? 0) + 1);
  }
  if (uniqueNums.size === 1) {
    return 0;
  }
  let res = 1;
  for (const [key, val] of map) {
    if (val > 1) {
      const splitedKeys = key.split(",");
      if (uniqueNums.has(Number(splitedKeys[0])) && uniqueNums.has(Number(splitedKeys[1]))) {
        res = Math.max(res, val);
      }
    }
  }
  if (res === 1) {
    return 0;
  }
  return res;
};

// sort + sliding window
function findLHS2(nums: number[]): number {
  nums.sort((a, b) => a - b);
  let j = 0;
  let res = 0;

  for (let i = 0; i < nums.length; i++) {
    while (nums[i] - nums[j] > 1) {
      j++;
    }
    if (nums[i] - nums[j] === 1) {
      res = Math.max(res, i - j + 1);
    }
  }
  return res;
};

