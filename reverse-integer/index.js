/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
  const max = Math.pow(2, 31);
  const isLessThanZero = x < 0;
  if (isLessThanZero) {
    x = -x;
  }
  const str = String(x);
  let res = 0;

  for (let i = str.length - 1; i >= 0; i--) {
    res += Number(str[i]) * Math.pow(10, i);
    if (res >= max) {
      return 0;
    }
  }
  if (isLessThanZero) {
    return -res;
  }
  return res;
};

console.log(reverse(-123));
