/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function (digits) {
  let needToPlus = true;

  for (let i = digits.length - 1; i >= 0; i--) {
    if (needToPlus) {
      digits[i]++;
    }
    if (digits[i] === 10) {
      digits[i] = 0;
      needToPlus = true;
    } else {
      needToPlus = false;
    }
  }
  if (digits[0] === 0) {
    return [1, ...digits];
  }
  return digits;
};

console.log(plusOne([1, 2, 3]));
console.log(plusOne([9, 9, 9]));
