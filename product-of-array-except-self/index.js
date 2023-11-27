/**
 * @param {number[]} nums
 * @return {number[]}
 */
const productExceptSelf = function (nums) {
  const numsLen = nums.length;
  const productPrefix = Array(numsLen);
  const productPostfix = Array(numsLen);
  const result = Array(numsLen);

  // calculate product prefix and postfix
  for (let i = 0; i < numsLen; i++) {
    if (i === 0) {
      productPrefix[i] = nums[i];
      productPostfix[numsLen - 1] = nums[numsLen - 1];
    } else {
      productPrefix[i] = productPrefix[i - 1] * nums[i];
      productPostfix[numsLen - 1 - i] =
        productPostfix[numsLen - i] * nums[numsLen - 1 - i];
    }
  }
  // calculate product result
  for (let j = 0; j < numsLen; j++) {
    if (j === 0) {
      result[j] = productPostfix[j + 1];
    } else if (j === numsLen - 1) {
      result[j] = productPrefix[j - 1];
    } else {
      result[j] = productPrefix[j - 1] * productPostfix[j + 1];
    }
  }
  return result;
};

console.log(productExceptSelf([1, 2, 3, 4]));
