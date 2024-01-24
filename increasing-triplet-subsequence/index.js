/**
 * @param {number[]} nums
 * @return {boolean}
 */
var increasingTriplet = function (nums) {
  let firstNum = Infinity;
  let secondNum = Infinity;

  for (const n of nums) {
    if (n <= firstNum) {
      firstNum = n;
    } else if (n <= secondNum) {
      secondNum = n;
    } else {
      return true;
    }
  }
  return false;
};
