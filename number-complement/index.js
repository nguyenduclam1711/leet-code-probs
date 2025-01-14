/**
 * @param {number} num
 * @return {number}
 */
var findComplement = function (num) {
  let temp = num;
  let mask = 0;

  while (temp > 0) {
    mask = (mask << 1) | 1;
    temp = temp >> 1;
  }
  return num ^ mask;
};
