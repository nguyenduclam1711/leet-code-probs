function makeKey(nums: number[]) {
  return nums.join(",");
}

function combinationSum(candidates: number[], target: number): number[][] {
  candidates.sort((a, b) => a - b);
  const dp: number[][][] = Array.from(Array(target + 1)).fill(new Array(0));

  for (let i = 2; i < dp.length; i++) {
    dp[i] = [];
    const c = new Map<string, boolean>();

    for (const num of candidates) {
      if (i < num) {
        break;
      } else if (i === num) {
        dp[i].push([num]);
      } else {
        const remain = i - num;
        for (const nums of dp[remain]) {
          const newCombination = [num, ...nums];
          newCombination.sort((a, b) => a - b);
          const key = makeKey(newCombination);
          if (c.has(key)) {
            continue;
          } else {
            c.set(key, true);
            dp[i].push(newCombination);
          }
        }
      }
    }
  }
  return dp[target];
};
