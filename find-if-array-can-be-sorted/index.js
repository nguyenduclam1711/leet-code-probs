function countSetBitsRecursive(n) {
  const remainder = n % 2;
  const divideVal = parseInt(n / 2);
  let res = 0;

  if (remainder === 1) {
    res++;
  }
  if (divideVal === 0) {
    return res;
  }
  res += countSetBitsRecursive(divideVal);
  return res;
}

function canSortArray(nums) {
  let currSetBits = countSetBitsRecursive(nums[0]);
  [prevMax, currMin, currMax] = [-Infinity, nums[0], nums[0]];

  for (let i = 1; i < nums.length; i++) {
    const n = nums[i];
    const setBits = countSetBitsRecursive(n);
    if (setBits === currSetBits) {
      if (n > currMax) {
        currMax = n;
      }
      if (n < currMin) {
        currMin = n;
      }
    } else {
      prevMax = currMax;
      [currMin, currMax] = [n, n];
      currSetBits = setBits;
    }
    if (currMin < prevMax) {
      return false;
    }
  }
  return true;
}
