/**
 * @param {number} n - a positive integer
 * @return {number}
 */
var hammingWeight = function (num) {
  let w = 0;

  while (num !== 0) {
    if (num & (1 === 1)) {
      w++;
    }
    num = num >>> 1;
  }
  return w;
};

const num = 128;
console.log(hammingWeight(num));
